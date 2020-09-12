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

	"github.com/abaskin/signald-go/signald"
	"github.com/spf13/cobra"
)

var (
	// username    string  see send.go
	timeOut    int
	returnJSON bool
)

// receiveCmd represents the receive command
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Query signald for new messages",
	Long:  `Query signald for new messages, with messages printed to standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan signald.RawResponse)
		sc := make(chan struct{})
		s.Receive(c, sc, username, timeOut, returnJSON)

		message := signald.RawResponse{}
		for {
			message = <-c

			jsonData, _ := json.Marshal(message)
			fmt.Println(string(jsonData))

			if message.Done {
				return
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(receiveCmd)

	receiveCmd.Flags().StringVarP(&username, "username", "u", "", "The username of the account to receive messages for")
	receiveCmd.MarkFlagRequired("username")
	receiveCmd.Flags().IntVarP(&timeOut, "timeOut", "t", 5, "Number of seconds to wait for new messages (0 to disable timeout), default is 5 seconds")
	receiveCmd.Flags().BoolVarP(&returnJSON, "jsonArray", "a", true, "Output received messages as a JSON array if timeout is set")
}
