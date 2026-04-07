package main

import (
	"fmt"
)

var AppVersion = "dev"

func main() {
	InitConfig()

	fmt.Printf("linear-cli %s\n", AppVersion)
}
