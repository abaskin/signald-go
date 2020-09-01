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

// addDevice Adds another device to a signal account that signald controls the master device on.
func (s *Signald) addDevice(username string, uri string) (Response, error) {
	if username == "" {
		return Response{}, s.MakeError("username is required")
	}

	if uri == "" {
		return Response{}, s.MakeError("uri is required")
	}

	return s.SendAndListen(Request{
		Type:     "add_device",
		Username: username,
		URI:      uri,
	}, []string{"device_added"})
}
