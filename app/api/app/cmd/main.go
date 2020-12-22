package main

import (
	"codebase/app/api/app/internal/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		panic(err)
	}
}
