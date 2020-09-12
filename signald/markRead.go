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

// MarkRead represents the listGroups command
func (s *Signald) MarkRead(username string, when int64, timeStamps []int64) (Response, error) {
	if username == "" {
		return Response{}, s.MakeError("username is required")
	}

	if len(timeStamps) == 0 {
		return Response{}, s.MakeError("timeStamps are required")
	}

	message, err := s.SendAndListen(
		Request{
			Type:       "",
			Username:   username,
			When:       when,
			Timestamps: timeStamps,
		},
		[]string{"marked_read", "untrusted_identity"})

	if err == nil && message.Type == "untrusted_identity" {
		err = s.MakeError("Untrusted Identity, see Response for details")
	}

	return message, err
}
