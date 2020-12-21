package main

import (
	"codebase/app/api/app/cmd/app"
)

func main() {
	if err := app.NewApiAppCommand().Execute(); err != nil {
		panic(err)
	}
}
