package cmd

import (
	"github.com/wemanity-belgium/hyperclair/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Create hyperclair Server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := server.ListenAndServe()

		return err
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}