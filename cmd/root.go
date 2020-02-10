package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "The web app",
	Long:  "The web server application",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
