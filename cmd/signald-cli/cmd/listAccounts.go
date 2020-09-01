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

// listAccountsCmd represents the listAccounts command
var listAccountsCmd = &cobra.Command{
	Use:   "listAccounts",
	Short: "List of all accounts registered to this signald instance",
	Long:  `Prints a list of all accounts registered to this signald instance`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.ListAccounts()

		accounts := []string{}
		for _, a := range message.Data.Accounts {
			accounts = append(accounts, fmt.Sprintf(
				"Username: %s DeviceID %d Filename %s Registered %t Subscribed %t HasKeys %t",
				a.Username, a.DeviceID, a.Filename, a.Registered, a.Subscribed, a.HasKeys))
		}

		handleReturn(message, err, strings.Join(accounts, "\n"))
	},
}

func init() {
	RootCmd.AddCommand(listAccountsCmd)
}
