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
	// username    string  see send.go
	code string
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a new number",
	Long:  `Completes the registration process, by providing a verification code sent after the register command. `,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.Verify(username, code)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(verifyCmd)

	verifyCmd.Flags().StringVarP(&username, "username", "u", "", "The username of the account")
	verifyCmd.MarkFlagRequired("username")
	verifyCmd.Flags().StringVarP(&code, "code", "c", "", "The verification code, the - in the middle code is optional")
	verifyCmd.MarkFlagRequired("code")
}
