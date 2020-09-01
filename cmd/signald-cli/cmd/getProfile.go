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

// getProfileCmd represents the listContacts command
var getProfileCmd = &cobra.Command{
	Use:   "getProfile",
	Short: "Gets a user's profile",
	Long:  `Gets a user's profile. At this time only the name is available. Must have the user's profileKey already, otherwise you'll get a profile_not_available error.`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.GetProfile(username, signald.RequestAddress{Number: recipientNumber})

		handleReturn(message, err, fmt.Sprintf(
			"Name: %s Avatar: %s IdentityKey: %s UnidentifiedAccess: %s UnrestrictedUnidentifiedAccess: %t",
			message.Data.Profile.Name, message.Data.Profile.Avatar, message.Data.Profile.IdentityKey,
			message.Data.Profile.UnidentifiedAccess, message.Data.Profile.UnrestrictedUnidentifiedAccess))
	},
}

func init() {
	RootCmd.AddCommand(getProfileCmd)

	getProfileCmd.Flags().StringVarP(&username, "username", "u", "", "The account to use to check the registration.")
	getProfileCmd.MarkFlagRequired("username")
	getProfileCmd.Flags().StringVarP(&recipientNumber, "recipientNumber", "r", "", "The full number to look up.")
	getProfileCmd.MarkFlagRequired("recipientNumber")
}
