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

// SetExpiration represents the set_expiration command
func (s *Signald) SetExpiration(username string, recipientAddress RequestAddress,
	recipientGroupID string, expiresInSeconds int) (Response, error) {
	if username == "" {
		return Response{}, s.MakeError("username is required")
	}

	recipientAddressEmpty := recipientAddress.Number == "" && recipientAddress.UUID == ""

	if (recipientAddressEmpty && recipientGroupID == "") ||
		(!recipientAddressEmpty && recipientGroupID != "") {
		return Response{}, s.MakeError("recipientNumber and recipientGroupId are mutually exclusive and one of them is required")
	}

	request := Request{
		Type:             "set_expiration",
		Username:         username,
		RecipientAddress: &recipientAddress,
		RecipientGroupID: recipientGroupID,
		ExpiresInSeconds: expiresInSeconds,
	}

	return s.SendAndListen(request, []string{"expiration_updated"})
}
