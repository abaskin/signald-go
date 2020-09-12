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

	"github.com/abaskin/signald-go/signald"
	"github.com/spf13/cobra"
)

// getUserCmd represents the listContacts command
var getUserCmd = &cobra.Command{
	Use:   "getUser",
	Short: "Checks whether a contact is currently registered with the server",
	Long:  `Checks whether a contact is currently registered with the server. Returns the contact's registration state.`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.GetUser(username, signald.RequestAddress{Number: recipientNumber})

		handleReturn(message, err, fmt.Sprintf("Token: %s Number: %s Voice: %t Video: %t",
			message.Data.ContactTokenDetails.Token, message.Data.ContactTokenDetails.Number,
			message.Data.ContactTokenDetails.Voice, message.Data.ContactTokenDetails.Video))
	},
}

func init() {
	RootCmd.AddCommand(getUserCmd)

	getUserCmd.Flags().StringVarP(&username, "username", "u", "", "The account to use to check the registration.")
	getUserCmd.MarkFlagRequired("username")
	getUserCmd.Flags().StringVarP(&recipientNumber, "recipientNumber", "r", "", "The full number to look up.")
	getUserCmd.MarkFlagRequired("recipientNumber")
}
