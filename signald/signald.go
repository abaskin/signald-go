// Copyright © 2018 Finn Herzfeld <finn@janky.solutions>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package signald

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"time"

	"github.com/mdp/qrterminal"
	"github.com/rs/xid"
)

// Signald is a connection to a signald instance.
type Signald struct {
	socket     net.Conn
	SocketPath string
	Verbose    bool
	StatusJSON bool
	LogJSON    []Response
}

// Connect connects to the signald socket
func (s *Signald) Connect() error {
	socket, err := net.Dial("unix", s.SocketPath)
	if err != nil {
		s.verbose("Connect Error: " + err.Error())
		return s.MakeError(err)
	}

	s.socket = socket
	s.verbose("Connected to signald socket " + s.socket.RemoteAddr().String())

	return nil
}

// Disconnect disconnects from the signald socket
func (s *Signald) Disconnect() error {
	socketPath := s.socket.RemoteAddr().String()
	if err := s.socket.Close(); err != nil {
		s.verbose("Disconnect Error: " + err.Error())
		return s.MakeError(err)
	}

	s.socket = nil
	s.verbose("Disconnected from signald socket" + socketPath)

	return nil
}

// Listen listens for events from signald
func (s *Signald) Listen(c chan RawResponse) {
	d := json.NewDecoder(s.socket)

	message := RawResponse{}

	for {
		if message.Error = d.Decode(&message.JSON); message.Error != nil {
			message.Error = s.MakeError(message.Error)
		}

		c <- message

		if message.Error != nil {
			return
		}
	}
}

// ListenFor listens for events from signald, stops when ID is found and
// populates the returned Response object as required
func (s *Signald) ListenFor(stopID string) (Response, error) {
	cs := make(chan RawResponse)
	go s.Listen(cs)

	for {
		msg := Response{}

		message := <-cs

		if message.Error != nil {
			return Response{}, s.MakeError(message.Error)
		}

		msg.ID = fmt.Sprintf("%s", message.JSON["id"])
		if msg.ID == stopID {
			msg.Type = fmt.Sprintf("%s", message.JSON["type"])

			msgData, haveData := message.JSON["data"]
			jsonData, _ := json.Marshal(msgData)

			switch msg.Type {
			case "send_results":
				json.Unmarshal(jsonData, &msg.Data.SendResults)

			case "user":
				json.Unmarshal(jsonData, &msg.Data.UserDetails)

			case "account_list":
				json.Unmarshal(jsonData, &msg.Data.Accounts)

			case "contact_list":
				json.Unmarshal(jsonData, &msg.Data.Contacts)

			case "group_list":
				json.Unmarshal(jsonData, &msg.Data.Groups)

			case "identities":
				json.Unmarshal(jsonData, &msg.Data.Identities)

			case "profile":
				json.Unmarshal(jsonData, &msg.Data.Profile)

			case
				"profile_not_available",
				"user_not_registered":
				return msg, s.MakeError(msg.Type)

			case "linking_uri":
				json.Unmarshal(jsonData, &msg.Data)
				b := bytes.NewBufferString(string(jsonData))
				if strings.HasPrefix(msg.ID, "false") {
					b.Reset()
					qrterminal.Generate(string(jsonData), qrterminal.M, b)
				}
				msg.Data.StatusMessage.Error = false
				msg.Data.StatusMessage.Message = b.String()

			case
				"verification_required",
				"verification_succeeded",
				"linking_successful":
				json.Unmarshal(jsonData, &msg.Data.Accounts[0])

			case
				"unexpected_error",
				"input_error",
				"trust_failed",
				"update_contact_error",
				"linking_error",
				"verification_error":
				json.Unmarshal(jsonData, &msg.Data.StatusMessage)
				return msg, s.MakeError(msg)

			case "error":
				msg.Data.StatusMessage.Error = true
				msg.Data.StatusMessage.Message = msgData.(string)
				return msg, s.MakeError(msg)

			default:
				if haveData {
					json.Unmarshal(jsonData, &msg.Data.StatusMessage)
				}
			}

			return msg, nil
		}
	}
}

// SendRequest sends a request to signald. Mostly used interally.
func (s *Signald) SendRequest(request Request) (string, error) {
	if request.ID == "" {
		request.ID = "signald-go-" + xid.New().String()
	}

	s.verbose("Request ID: " + request.ID)

	b, err := json.Marshal(request)
	if err == nil {
		s.verbose("Sending " + string(b))

		err = json.NewEncoder(s.socket).Encode(request)
	}

	return request.ID, err
}

// SendAndListen sends a request to signald, listens for the response and returns it
func (s *Signald) SendAndListen(request Request, success []string) (Response, error) {
	var err error

	defer func() {
		if r := recover(); r != nil && err != nil {
			err = s.MakeError(r)
		}
	}()

	if s.socket == nil {
		if err = s.Connect(); err != nil {
			err = s.MakeError(err)
			return Response{}, err
		}
		defer s.Disconnect()
	}

	requestID := ""
	requestID, err = s.SendRequest(request)
	if err != nil {
		err = s.MakeError(err)
		return Response{}, err
	}

	message := Response{}
	message, err = s.ListenFor(requestID)
	if err != nil {
		err = s.MakeError(err)
		return Response{}, err
	}

	for _, s := range success {
		if message.Type == s {
			return message, err
		}
	}

	err = s.MakeError(message)
	return message, err
}

// verbose print log message if Verbose is set taking into account JsonStatus
func (s *Signald) verbose(logMsg string) {
	if !s.Verbose {
		return
	}

	if !s.StatusJSON {
		log.Print(logMsg)
		return
	}

	s.LogJSON = append(s.LogJSON, Response{
		Type: "log",
		Data: ResponseData{
			Timestamp: time.Now().Format(time.RFC3339),
			StatusMessage: StatusMessage{
				Message: logMsg,
			},
		},
	})
}

// MakeError make and return an error with the function name
func (s *Signald) MakeError(err interface{}) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	var errFormat string
	switch err.(type) {
	case string:
		errFormat = "%s"

	case error:
		errFormat = "%w"

	case Response:
		errFormat = "%s"
		err = fmt.Sprintf("%s: %s", err.(Response).Type, err.(Response).Data.StatusMessage.Message)

	default:
		errFormat = "%+v"
	}

	return fmt.Errorf("%s:%d - "+errFormat, frame.Function, frame.Line, err)
}
