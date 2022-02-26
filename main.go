package main

import (
	"os"
	"rings/cmd"
)

func main() {
	if err := cmd.Cmds().Execute(); err != nil {
		os.Exit(1)
	}
}


