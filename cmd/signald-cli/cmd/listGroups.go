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
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listGroupsCmd represents the listGroups command
var listGroupsCmd = &cobra.Command{
	Use:   "listGroups",
	Short: "List of all the groups of which the user is a member",
	Long:  `Prints a list of all groups of which user is a member`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.ListGroups(username)

		active := false
		groups := []string{}
		for _, group := range message.Data.Groups {
			members := [][]string{}
			for _, m := range group.Members {
				members[0] = append(members[0], m.Number)
				if username == m.Number {
					active = true
				}
			}

			b := new(bytes.Buffer)
			csv.NewWriter(b).WriteAll(members)

			groups = append(groups,
				fmt.Sprintf("Id: %s Name: %s Active: %t Blocked: %t Members: %s",
					group.GroupID, group.Name, active, false, b.String()))
		}

		handleReturn(message, err, strings.Join(groups, "\n"))
	},
}

func init() {
	RootCmd.AddCommand(listGroupsCmd)

	listGroupsCmd.Flags().StringVarP(&username, "username", "u", "", "The username of the account to use)")
	listGroupsCmd.MarkFlagRequired("username")
}
