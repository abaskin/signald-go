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

// Register represents the register command
func (s *Signald) Register(username string, captcha string, voice bool) (Response, error) {
	if username == "" {
		return Response{}, s.MakeError("username is required")
	}

	return s.SendAndListen(Request{
		Type:     "register",
		Username: username,
		Captcha:  captcha,
		Voice:    voice,
	}, []string{"verification_required"})
}
