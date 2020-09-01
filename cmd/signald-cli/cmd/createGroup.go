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
	"strings"

	"github.com/abaskin/signald-go/signald"
	"github.com/spf13/cobra"
)

var (
	// username    string  see send.go
	recipientGroupID string
	groupName        string
	members          string
	groupAvatar      string
)

// createGroupCmd represents the createGroup command
var createGroupCmd = &cobra.Command{
	Use:   "createGroup",
	Short: "Create or update a group",
	Long:  `Create or update a group`,
	Run: func(cmd *cobra.Command, args []string) {
		memberList := []string{}
		r := csv.NewReader(strings.NewReader(members))
		m, err := r.ReadAll()
		if err == nil {
			handleReturn(signald.Response{}, s.MakeError("Invalid member list: "+err.Error()), "")
		}
		memberList = m[0]

		message, err := s.CreateGroup(username, recipientGroupID, groupName,
			memberList, groupAvatar)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(createGroupCmd)

	createGroupCmd.Flags().StringVarP(&username, "username", "u", "", "The username of the account to create/update the group)")
	createGroupCmd.MarkFlagRequired("username")
	createGroupCmd.Flags().StringVarP(&recipientGroupID, "recipientGroupID", "r", "", "The base64 encoded group ID, if not specified a new group will be created")
	createGroupCmd.Flags().StringVarP(&groupName, "groupName", "g", "", "The value to which the group name is set")
	createGroupCmd.Flags().StringVarP(&members, "members", "m", "", "A list of users (full international format phone numbers) that should be added to the group")
	createGroupCmd.Flags().StringVarP(&groupAvatar, "groupAvatar", "a", "", "The avatar to set as the group's avatar (The actual format is unknown, probably a path to a file on the disk)")
}
