package main

import (
	"os"

	"github.com/abdulkareem83/sqlmetricsbeat/cmd"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/abdulkareem83/sqlmetricsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
