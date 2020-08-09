package app

import (
	"codebase/app/api/app/bootstrap"

	"github.com/spf13/cobra"
)

func NewStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start api server",
		Long:  "Start the api application",
		Run: func(cmd *cobra.Command, args []string) {
			cfgFile, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			bootstrap.Start(cfgFile, cmd.Flags())
		},
	}
	cmd.Flags().StringP("config", "c", "config.yaml", "app config file")
	return cmd
}
