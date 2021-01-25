package cmd

import (
	"codebase/app/api/internal/bootstrap"
	"codebase/app/api/internal/cfg"
	"codebase/pkg/app"
	"codebase/pkg/defers"
	"log"

	"go.uber.org/zap"

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

			defer defers.Run()

			//初始化全局变量，资源等
			application, err := app.New(cfg.Config, cfgFile, cmd.Flags())
			if err != nil {
				panic(err)
			}

			if err = application.RunWith(bootstrap.Start); err != nil {
				log.Panic("application run error", zap.Error(err))
			}
		},
	}
	cmd.Flags().StringP("config", "c", "config.yaml", "app config file")
	return cmd
}
