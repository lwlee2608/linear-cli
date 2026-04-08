package main

import "os"

var AppVersion = "dev"

func main() {
	rootCmd.Version = AppVersion
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
