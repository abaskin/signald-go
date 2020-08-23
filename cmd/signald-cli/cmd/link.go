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
	"log"
	"math/rand"
	"os"

	"github.com/mdp/qrterminal"
	"github.com/spf13/cobra"

	"github.com/abaskin/signald-go/signald"
)

var uriOrQR bool

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link to an existing Signal account",
	Long:  `Get a URI or QR code to link to an existing Signal account`,
	Run: func(cmd *cobra.Command, args []string) {
		requestID := fmt.Sprint("signald-cli-", rand.Intn(1000))
		s.SendRequest(signald.Request{
			Type: "link",
			ID:   requestID,
		})

		c := make(chan signald.Response)
		go s.Listen(c)
		for {
			message := <-c
			if message.ID == requestID {
				switch message.Type {
				case "linking_error":
					log.Fatal(message.Data.Message)
					os.Exit(1)
					break

				case "linking_uri":
					if uriOrQR {
						fmt.Println(message.Data.URI)
					} else {
						qrterminal.Generate(message.Data.URI, qrterminal.M, os.Stdout)
					}
					break

				case "linking_successful":
					if !uriOrQR {
						fmt.Println("Successfully linked")
						os.Exit(0)
					}
					break
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(linkCmd)
	linkCmd.Flags().BoolVarP(&uriOrQR, "uri", "u", false, "Print a URI instead of a QR code.")
}
