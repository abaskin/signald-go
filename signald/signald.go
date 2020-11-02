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
	"fmt"
	"log"
	"net"
	"runtime"
	"time"

	jsoniter "github.com/json-iterator/go"
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

// IsConnected check to see if the socket is connected
func (s *Signald) IsConnected() bool {
	return s.socket != nil
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
	if err := s.socket.Close(); err != nil {
		s.verbose("Disconnect Error: " + err.Error())
		return s.MakeError(err)
	}

	s.socket = nil
	s.verbose("Disconnected from signald socket" + s.SocketPath)

	return nil
}

// Listen listens for events from signald
func (s *Signald) Listen(c chan RawResponse) {
	d := jsoniter.NewDecoder(s.socket)

	message := RawResponse{}

	for {
		message.Error = s.MakeError(d.Decode(&message))

		c <- message

		if message.Error != nil {
			return
		}
	}
}

// ListenFor listens for events from signald and stops when the passed ID is found
func (s *Signald) ListenFor(stopID string) (Response, error) {
	cs := make(chan RawResponse)
	go s.Listen(cs)

	for {
		message := <-cs

		if message.Error != nil {
			return Response{}, s.MakeError(message.Error)
		}

		if message.ID == stopID {
			response := Response{}

			jsonData, _ := jsoniter.Marshal(message)
			jsoniter.Unmarshal(jsonData, &response)

			if response.Data.StatusMessage.Error {
				return response, s.MakeError(response)
			}

			return response, nil
		}
	}
}

// SendRequest sends a request to signald. Mostly used interally.
func (s *Signald) SendRequest(request Request) (string, error) {
	if request.ID == "" {
		request.ID = "signald-go-" + xid.New().String()
	}

	s.verbose("Request ID: " + request.ID)

	b, err := jsoniter.Marshal(request)
	if err == nil {
		s.verbose("Sending " + string(b))
		err = jsoniter.NewEncoder(s.socket).Encode(request)
	}

	return request.ID, err
}

// SendAndListen sends a request to signald, listens for the response and returns it.
// If the request type is empty only listen (used by link).
func (s *Signald) SendAndListen(request Request, success []string) (Response, error) {
	var err error

	defer func() {
		if r := recover(); r != nil && err == nil {
			err = s.MakeError(r)
		}
	}()

	if !s.IsConnected() {
		if err = s.Connect(); err != nil {
			err = s.MakeError(err)
			return Response{}, err
		}
		defer s.Disconnect()
	}

	if request.Type != "" {
		request.ID, err = s.SendRequest(request)
		if err != nil {
			err = s.MakeError(err)
			return Response{}, err
		}
	}

	response := Response{}
	response, err = s.ListenFor(request.ID)
	if err != nil {
		err = s.MakeError(err)
		return Response{}, err
	}

	for _, s := range success {
		if response.Type == s {
			return response, err
		}
	}

	err = s.MakeError(response)
	return response, err
}

// verbose print log message if Verbose is set taking into account JsonStatus
func (s *Signald) verbose(logMsg string) {
	if !s.Verbose {
		return
	}

	if !s.StatusJSON {
		log.Println(logMsg)
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
	if err == nil {
		return nil
	}

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
