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
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/abaskin/signald-go/signald"
	"github.com/spf13/cobra"
)

var (
	username    string
	toUser      string
	toGroup     string
	messageBody string
	attach      string
	quote       string
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to another user or group",
	Long:  `Send a message to another user or group on Signal`,
	Run: func(cmd *cobra.Command, args []string) {
		quoteObj := signald.RequestQuote{}
		if err := json.Unmarshal([]byte(quote), &quoteObj); err != nil {
			err = fmt.Errorf("Invalid Quote JSON: %w", err)
			handleReturn(signald.Response{}, err, "")
		}

		attachments := []signald.RequestAttachment{}
		attach = strings.TrimSpace(attach)
		if strings.HasPrefix(attach, "[") && strings.HasSuffix(attach, "]") {
			if err := json.Unmarshal([]byte(attach), &attachments); err != nil {
				err = fmt.Errorf("Invalid Attachment JSON: %w", err)
				handleReturn(signald.Response{}, err, "")
			}
		}

		if attach != "" && len(attachments) == 0 {
			af, err := csv.NewReader(strings.NewReader(attach)).ReadAll()
			if err != nil {
				err = fmt.Errorf("Invalid Attachment File Path List: %w", err)
				handleReturn(signald.Response{}, err, "")
			}
			for _, fileName := range af[0] {
				attachments = append(attachments, signald.RequestAttachment{
					Filename: fileName,
				})
			}
		}

		message, err := s.Send(username, signald.RequestAddress{Number: toUser},
			toGroup, messageBody, attachments, quoteObj)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&username, "username", "u", "", "The username to send from (required)")
	sendCmd.MarkFlagRequired("username")
	sendCmd.Flags().StringVarP(&toUser, "to", "t", "", "The user to send the message to (cannot be combined with --group)")
	sendCmd.Flags().StringVarP(&toGroup, "group", "g", "", "The group to send the message to (cannot be combined with --to)")
	sendCmd.Flags().StringVarP(&messageBody, "message", "m", "", "The text of the message to send")
	sendCmd.Flags().StringVarP(&quote, "quote", "q", "{}", "A vaid signald quote JSON object, see the signald documation for the details")
	sendCmd.Flags().StringVarP(&attach, "attachments", "a", "", `A list of file paths to attach to the message.
	This can also be a valid JASON array (see below), only Filename is required.
	[{
		"Filename": "/tmp/file.bmp",
		"Caption": "Some Caption",
		"Width": 100,
		"Height": 100,
		"VoiceNote": false,
		"Preview": true
	}]`)
}
