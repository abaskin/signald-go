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

// listGroupsCmd represents the listGroups command
var listGroupsCmd = &cobra.Command{
	Use:   "listGroups",
	Short: "list of all the groups that the user is in.",
	Long:  `Prints a list of all groups the user is in to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		requestID := fmt.Sprint("signald-cli-", rand.Intn(1000))
		s.SendRequest(signald.Request{
			Type:     "list_groups",
			Username: username,
			ID:       requestID,
		})

		c := make(chan signald.Response)
		go s.Listen(c)
		for {
			message := <-c
			if message.ID == requestID {
				for _, group := range message.Data.Groups {
					fmt.Println(group.Name)
				}
				break
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listGroupsCmd)

	listGroupsCmd.Flags().StringVarP(&username, "username", "u", "", "The username of the account to use)")
	listGroupsCmd.MarkFlagRequired("username")
}
