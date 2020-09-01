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

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/abaskin/signald-go/signald"
)

// handlerReturn handle the results of a command
func handleReturn(message signald.Response, err error, customOut string) {
	if err != nil {
		if statusJSON {
			s.LogJSON = append(s.LogJSON, signald.Response{
				Type: "error",
				Data: signald.ResponseData{
					StatusMessage: message.Data.StatusMessage,
					Timestamp:     time.Now().Format(time.RFC3339),
				},
			})
			message.Type = ""
		} else {
			customOut = fmt.Sprintf("Error: %+v", err)
		}
		defer os.Exit(1)
	}

	if statusJSON {
		if message.Type != "" {
			s.LogJSON = append(s.LogJSON, message)
		}
		jsonData, _ := json.Marshal(s.LogJSON)
		customOut = string(jsonData)
	}

	if customOut == "" {
		customOut = fmt.Sprintf("%s: %s", message.Type, message.Data.StatusMessage.Message)
	}

	fmt.Println(customOut)
}
