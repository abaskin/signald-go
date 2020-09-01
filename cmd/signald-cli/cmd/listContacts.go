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
	"strings"

	"github.com/spf13/cobra"
)

// listContactsCmd represents the listContacts command
var listContactsCmd = &cobra.Command{
	Use:   "listContacts",
	Short: "Lists all of the contacts for the specified user",
	Long:  `Lists all of the contacts in the contact store for the specified user`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.ListContacts(username)

		contacts := []string{}
		for _, c := range message.Data.Contacts {
			contacts = append(contacts, fmt.Sprintf(
				"Number %s ProfileKey %s MessageExpirationTime %d InboxPosition %d Address.Number %s Address.UUID %s",
				c.Name, c.ProfileKey, c.MessageExpirationTime, c.MessageExpirationTime,
				c.Address.Number, c.Address.UUID))
		}

		handleReturn(message, err, strings.Join(contacts, "\n"))
	},
}

func init() {
	RootCmd.AddCommand(listContactsCmd)

	listContactsCmd.Flags().StringVarP(&username, "username", "u", "", "The username of the account for which to list contacts")
	listContactsCmd.MarkFlagRequired("username")
}
