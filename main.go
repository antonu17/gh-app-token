package main

import (
	"os"

	"github.com/n26/gh-app-token/cmd"
)

func main() {
	if err := cmd.Execute(os.Args[1:], os.Stdout); err != nil {
		os.Exit(1)
	}
}
