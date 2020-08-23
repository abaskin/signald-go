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
	"log"

	"github.com/spf13/cobra"

	"github.com/abaskin/signald-go/signald"
)

var (
	username    string
	toUser      string
	toGroup     string
	messageBody string
	attachment  string
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send a message to another user or group",
	Long:  `send a message to another user or group on Signal`,
	Run: func(cmd *cobra.Command, args []string) {
		request := signald.Request{
			Type:     "send",
			Username: username,
		}

		if toUser != "" {
			request.RecipientNumber = toUser
		} else if toGroup != "" {
			request.RecipientGroupID = toGroup
		} else {
			log.Fatal("--to or --group must be specified!")
		}

		if messageBody != "" {
			request.MessageBody = messageBody
		}

		if attachment != "" {
			request.AttachmentFilenames = []string{attachment}
		}
		s.SendRequest(request)
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&username, "username", "u", "", "The username to send from (required)")
	sendCmd.MarkFlagRequired("username")

	sendCmd.Flags().StringVarP(&toUser, "to", "t", "", "The user to send the message to (cannot be combined with --group)")

	sendCmd.Flags().StringVarP(&toGroup, "group", "g", "", "The group to send the message to (cannot be combined with --to)")

	sendCmd.Flags().StringVarP(&messageBody, "message", "m", "", "The text of the message to send")

	sendCmd.Flags().StringVarP(&attachment, "attachment", "a", "", "A file to attach to the message")
}
