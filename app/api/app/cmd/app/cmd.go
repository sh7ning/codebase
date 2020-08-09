package app

import (
	"github.com/spf13/cobra"
)

func NewApiAppCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "The api app",
		Long:  "The api application is a daemon program that serves all api requests",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				panic(err)
			}
		},
	}
	cmd.AddCommand(NewStartCommand())
	return cmd
}
