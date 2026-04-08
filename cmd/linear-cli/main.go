package main

import (
	"fmt"
	"os"
)

var AppVersion = "dev"

func main() {
	rootCmd.Version = AppVersion
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
