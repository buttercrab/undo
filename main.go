package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func printInit() {
	fmt.Println("export UNDO_INIT=true")
	fmt.Println("mkdir -p ~/.undo")
	os.Exit(0)
}

func checkInit() (bool, error) {
	result, err := exec.Command("echo", "$UNDO_INIT").Output()
	if err != nil {
		return false, err
	}
	return string(result) == "true", nil
}

func main() {
	initCheck, err := checkInit()
	if err != nil {
		log.Fatal(err)
	}
	initHelpMessage := "Inits the undo command by 'eval $(undo --init)'"
	if !initCheck {
		initHelpMessage += "\nYou terminal doesn't seem to be initialized.\nPut 'eval $(undo --init)' in your ~/.bashrc or ~/.zshrc"
	}
	initPtr := flag.Bool("init", false, initHelpMessage)
	flag.Parse()

	if *initPtr {
		printInit()
	}
}
