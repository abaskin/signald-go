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

	"github.com/abaskin/signald-go/signald"
	"github.com/spf13/cobra"
)

// getIdentitiesCmd represents the listContacts command
var getIdentitiesCmd = &cobra.Command{
	Use:     "listIdentities",
	Aliases: []string{"getIdentities"},
	Short:   "Returns all known identities/keys for a specific number",
	Long:    `Returns all known identities/keys for a specific number`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.ListIdentities(username, signald.RequestAddress{Number: recipientNumber})

		identities := []string{}
		for _, i := range message.Data.Identities {
			identities = append(identities,
				fmt.Sprintf("Username: %s Added: %d Fingerprint: %s TrustLevel: %s SafetyNumber: %s",
					i.Address.Number, i.Added, i.Fingerprint, i.TrustLevel, i.SafetyNumber))
		}

		handleReturn(message, err, strings.Join(identities, "\n"))
	},
}

func init() {
	RootCmd.AddCommand(getIdentitiesCmd)

	getIdentitiesCmd.Flags().StringVarP(&username, "username", "u", "", "The local account to use to check the identity")
	getIdentitiesCmd.MarkFlagRequired("username")
	getIdentitiesCmd.Flags().StringVarP(&recipientNumber, "recipientNumber", "r", "", "The full number to look up")
}
