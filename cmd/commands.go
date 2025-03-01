package cmd

import (
	"log"
	"week/internal/cli"
)

func Execute() {

	rootCmd := cli.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
