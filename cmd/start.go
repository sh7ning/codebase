package cmd

import (
	"app/pkg/bootstrap"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start app",
	Long:  "Start the web server application",
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, err := cmd.Flags().GetString("config")
		if err != nil {
			panic(err)
		}

		bootstrap.Run(cfgFile)
	},
}

func init() {
	startCmd.Flags().StringP("config", "c", "config.yaml", "app config file")
}
