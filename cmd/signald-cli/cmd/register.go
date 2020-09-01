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

var (
	// username string
	captcha string
	voice   bool
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new number",
	Long:  `Begins the process of registering a new number on signal for use with signald`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.Register(username, captcha, voice)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&username, "username", "u", "", "The username of the account")
	registerCmd.MarkFlagRequired("username")
	registerCmd.Flags().StringVarP(&captcha, "captcha", "c", "", "The captcha value to use.")
	registerCmd.Flags().BoolVarP(&voice, "phone", "p", false, "Indicates if the verification code should be sent via a phone call. If false or not set the verification is done via SMS")
}
