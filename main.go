package main

import (
	"flag"
	"fmt"
	"os"
)

func printInit() {
	fmt.Print("export UNDO_INIT=true;")
	fmt.Print("mkdir -p ~/.undo;")
	os.Exit(0)
}

func main() {
	initHelpMessage := "Inits the undo command by 'eval $(undo --init)'"
	if os.Getenv("UNDO_INIT") != "true" {
		initHelpMessage += "\nYou terminal doesn't seem to be initialized.\nPut 'eval $(undo --init)' in your ~/.bashrc or ~/.zshrc"
	}
	initPtr := flag.Bool("init", false, initHelpMessage)
	flag.Parse()

	if *initPtr {
		printInit()
	}
}
