package main

import (
	"codebase/app/api/app/cmd/app"
)

func main() {
	cmd := app.NewApiAppCommand()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
