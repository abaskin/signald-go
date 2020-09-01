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
	"github.com/abaskin/signald-go/signald"
	"github.com/spf13/cobra"
)

var (
	// username    string
	recipientNumber string
	// recipientGroupID string
	expiresInSeconds int
)

// setExpirationCmd represents the set_expiration command
var setExpirationCmd = &cobra.Command{
	Use:   "setExpiration",
	Short: "Sets or changes the expiration time for messages in a group or PM",
	Long:  `Sets or changes the expiration time for messages in a group or PM, recipientNumber and recipientGroupId are mutually exclusive and one of them is required.`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.SetExpiration(username, signald.RequestAddress{Number: recipientNumber},
			recipientGroupID, expiresInSeconds)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(setExpirationCmd)

	setExpirationCmd.Flags().StringVarP(&username, "username", "u", "", "The username to send from (required)")
	setExpirationCmd.MarkFlagRequired("username")
	setExpirationCmd.Flags().StringVarP(&recipientNumber, "recipientNumber", "n", "", "The PM to change expiration for (cannot be combined with --recipientGroupID)")
	setExpirationCmd.Flags().StringVarP(&recipientGroupID, "recipientGroupID", "g", "", "The group ID to update expiration for (cannot be combined with --recipientNumber)")
	setExpirationCmd.Flags().IntVarP(&expiresInSeconds, "expiresInSeconds", "e", 0, "The number of seconds after which messages in the conversation should expire. Set to 0 to turn off disappearing messages.")
}
