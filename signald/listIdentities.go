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

// ListIdentities represents the get_identities command
func (s *Signald) ListIdentities(username string, recipientAddress RequestAddress) (Response, error) {
	if username == "" {
		return Response{}, s.MakeError("username is required")
	}

	request := Request{
		Type:     "get_identities",
		Username: username,
	}

	if !recipientAddress.Empty() {
		request.RecipientAddress = &recipientAddress
	}

	message, err := s.SendAndListen(request, []string{"identities"})

	return message, err
}
