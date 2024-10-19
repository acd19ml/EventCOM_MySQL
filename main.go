package main

import (
	"fmt"

	"github.com/acd19ml/EventCOM_MySQL/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
