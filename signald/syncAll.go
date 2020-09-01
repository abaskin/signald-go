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

// SyncAll sync contacts, groups and configuration
func (s *Signald) SyncAll(username string) (Response, error) {
	if username == "" {
		return Response{}, s.MakeError("username is required")
	}

	var message Response
	var err error
	for _, t := range []string{"sync_contacts", "sync_groups", "sync_configuration"} {
		message, err = s.SendAndListen(Request{
			Type:     t,
			Username: username,
		}, []string{"sync_requested"})
		if err != nil {
			return message, err
		}
	}

	return message, err
}
