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
	"log"
	"net"
)

// Signald is a connection to a signald instance.
type Signald struct {
	socket     net.Conn
	SocketPath string
}

func crash(err error) {
	if err != nil {
		panic(err)
	}
}

// Connect connects to the signad socket
func (s *Signald) Connect() error {
	if s.SocketPath == "" {
		s.SocketPath = "/var/run/signald/signald.sock"
	}
	socket, err := net.Dial("unix", s.SocketPath)
	if err != nil {
		return err
	}
	s.socket = socket
	log.Print("Connected to signald socket ", socket.RemoteAddr().String())
	return nil
}

// Listen listens for events from signald
func (s *Signald) Listen(c chan Response) error {
	// we create a decoder that reads directly from the socket
	d := json.NewDecoder(s.socket)

	var msg Response

	for {
		if err := d.Decode(&msg); err != nil {
			return err
		}
		c <- msg
	}
}

// SendRequest sends a request to signald. Mostly used interally.
func (s *Signald) SendRequest(request Request) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	log.Print("Sending ", string(b))
	e := json.NewEncoder(s.socket)
	return e.Encode(request)
}
