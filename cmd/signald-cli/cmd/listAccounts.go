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
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"

	"github.com/abaskin/signald-go/signald"
)

// listAccountsCmd represents the listAccounts command
var listAccountsCmd = &cobra.Command{
	Use:   "listAccounts",
	Short: "list of all the accounts registered to this signald instance.",
	Long:  `Prints a list of all users to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		requestID := fmt.Sprint("signald-cli-", rand.Intn(1000))
		s.SendRequest(signald.Request{
			Type: "list_accounts",
			ID:   requestID,
		})

		c := make(chan signald.Response)
		go s.Listen(c)
		for {
			message := <-c
			if message.ID == requestID {
				for _, account := range message.Data.Accounts {
					fmt.Println(account.Username)
				}
				break
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listAccountsCmd)
}
