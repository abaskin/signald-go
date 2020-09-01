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
	uri string
)

// addDeviceCmd represents the add_device command
var addDeviceCmd = &cobra.Command{
	Use:   "addDevice",
	Short: "Adds another device to a signal account that signald controls the master device",
	Long:  `Adds another device to a signal account that signald controls the master device`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.Verify(username, uri)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(addDeviceCmd)

	addDeviceCmd.Flags().StringVarP(&username, "username", "u", "", "The account to which to add the device")
	addDeviceCmd.MarkFlagRequired("username")
	addDeviceCmd.Flags().StringVarP(&uri, "uri", "i", "", "The tsdevice: URI that is provided by the other device (displayed as a QR code normally)")
	addDeviceCmd.MarkFlagRequired("uri")
}
