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

// Send represents the send command
func (s *Signald) Send(username string, toAddress RequestAddress, toGroup string,
	messageBody string, attachments []RequestAttachment, quote RequestQuote) (Response, error) {
	if username == "" {
		return Response{}, s.MakeError("username is required")
	}

	toAddressEmpty := toAddress.Number == "" && toAddress.UUID == ""

	if toAddressEmpty && toGroup == "" {
		return Response{}, s.MakeError("toUser or toGroup must be specified")
	}

	if !toAddressEmpty && toGroup != "" {
		return Response{}, s.MakeError("toUser and toGroup may not both be specified")
	}

	request := Request{
		Type:             "send",
		Username:         username,
		RecipientGroupID: toGroup,
		MessageBody:      messageBody,
		Attachments:      attachments,
	}

	if !toAddressEmpty {
		request.RecipientAddress = &toAddress
	}

	if quote.ID != 0 {
		request.Quote = &quote
	}

	return s.SendAndListen(request, []string{"send_results"})
}
