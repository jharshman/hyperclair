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

	"github.com/spf13/cobra"

	"github.com/jgsqware/hyperclair/ping"
	"strconv"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the registry",
	Long: `Ping the Docker registry to check if it's up`,
	Run: func(cmd *cobra.Command, args []string) {
		clairURI := RootCmd.PersistentFlags().Lookup("clair_uri").Value.String()
		clairPort, _:= strconv.Atoi(RootCmd.PersistentFlags().Lookup("clair_port").Value.String())
		registryURI := RootCmd.PersistentFlags().Lookup("registry_uri").Value.String()
		registryPort, _:= strconv.Atoi(RootCmd.PersistentFlags().Lookup("registry_port").Value.String())
		//TODO the Get Value is not great
		ping.Clair(clairURI,clairPort)
		ping.Registry(registryURI,registryPort)
	},
}

func init() {
	RootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}