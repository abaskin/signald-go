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
	"github.com/spf13/cobra"
)

// setProfileCmd represents the listContacts command
var setProfileCmd = &cobra.Command{
	Use:   "setProfile",
	Short: "Sets the user's profile, at this time only the name is available",
	Long:  `Sets the user's profile, at this time only the name is available`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.SetProfile(username, name)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(setProfileCmd)

	setProfileCmd.Flags().StringVarP(&username, "username", "u", "", "The account to use to set the profile")
	setProfileCmd.MarkFlagRequired("username")
	setProfileCmd.Flags().StringVarP(&name, "name", "n", "", "The name to set")
	setProfileCmd.MarkFlagRequired("name")
}
