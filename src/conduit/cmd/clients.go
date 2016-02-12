// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"conduit/log"
	"github.com/spf13/cobra"
)

// clientsCmd represents the clients command
var clientsCmd = &cobra.Command{
	Use:   "clients",
	Short: "Get the status of all mailboxes.",
	Long:  `Returns the connection status of all mailboxes.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ClientFromConfig()
		if err != nil {
			log.Debug(err.Error())
			log.Fatal("Could not configure client")
		}
		stats, err := client.ClientStatus()
		if err != nil {
			log.Debug(err.Error())
			log.Fatal("Could not retrieve statistics")
		}
		for mb, v := range stats {
			var vStr = ""
			if v {
				vStr = "ONLINE "
			} else {
				vStr = "OFFLINE"
			}
			if cmd.Flag("offline").Value.String() == "false" || !v {
				log.Status(mb, vStr, v)
			}
		}
	},
}

func init() {
	statsCmd.AddCommand(clientsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	clientsCmd.Flags().BoolP("offline", "x", false, "Show only offline clients")
}
