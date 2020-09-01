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
	uriOrQR    bool
	deviceName string
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link to an existing Signal account",
	Long:  `Get a URI or QR code to link to an existing Signal account`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := s.Link(deviceName, uriOrQR)

		handleReturn(message, err, "")
	},
}

func init() {
	RootCmd.AddCommand(linkCmd)

	linkCmd.Flags().StringVarP(&deviceName, "deviceName", "d", "", "The device name")
	linkCmd.MarkFlagRequired("deviceName")
	linkCmd.Flags().BoolVarP(&uriOrQR, "uri", "u", false, "Print a URI instead of a QR code")
}
