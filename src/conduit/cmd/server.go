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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"postmaster/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Conduit in server mode.",
	Long: `Run the conduit message server. To manage the server use the server sub
commands. For help run 'conduit help server'.`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.IsSet("enable_long_polling") {
			server.EnableLongPolling = viper.GetBool("enable_long_polling")
		}
		log.LogFile = true
		err := server.Start(viper.GetString("host"))
		fmt.Println("Could not start server:", err)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
