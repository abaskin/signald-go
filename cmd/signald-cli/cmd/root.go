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
	"os"

	"github.com/spf13/cobra"

	"github.com/abaskin/signald-go/signald"
)

var (
	socketPath string
	statusJSON bool
	verbose    bool
	s          *signald.Signald
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "signald-cli",
	Short: "Interact with a running signald instance",
	Long:  `signald-cli is a command line tool to interact with signald.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		s = &signald.Signald{
			SocketPath: socketPath,
			Verbose:    verbose,
			StatusJSON: statusJSON,
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&socketPath, "socket", "s", "/var/run/signald/signald.sock", "the path to the signald socket file")
	RootCmd.PersistentFlags().BoolVarP(&statusJSON, "json", "j", false, "return the results of the command as a JSON array")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "provide verbose logging")
}
