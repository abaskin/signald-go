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
	"encoding/json"
	"time"
)

// Receive represents the receive command
// The socket needs to be closed to stop the goroutine that reads from it if you
// are doing your own socket management.
func (s *Signald) Receive(c chan RawResponse, stopC chan struct{}, username string,
	timeOut int, returnJSON bool) {
	message := RawResponse{}

	if username == "" {
		message.Error = s.MakeError("username is required")
		c <- message
		return
	}

	if s.socket == nil {
		if message.Error = s.Connect(); message.Error != nil {
			c <- message
			return
		}
		defer s.Disconnect()
	}

	s.Subscribe(username)
	defer s.Unsubscribe(username)

	done := false
	if timeOut > 0 {
		go func() {
			<-time.After(time.Duration(timeOut) * time.Second)
			done = true
		}()
	}

	jsonSlice := []map[string]interface{}{}
	cs := make(chan RawResponse)
	go s.Listen(cs)
	for {
		select {
		case <-stopC:
			done = true

		case message = <-cs:
			if message.Error != nil {
				done = true
			}

			if returnJSON && timeOut != 0 {
				jsonSlice = append(jsonSlice, message.JSON)
				continue
			}

			jsonData, _ := json.Marshal(message.JSON)
			message.Data = string(jsonData)

			c <- message

		default:
			if done {
				message = RawResponse{
					Done: true,
					JSON: map[string]interface{}{
						"Data": jsonSlice,
					},
				}

				jsonData, _ := json.Marshal(message.JSON)
				message.Data = string(jsonData)

				c <- message

				return
			}
		}
	}
}
