package main

import (
	"github.com/buttercrab/undo/util"
	flag "github.com/integrii/flaggy"
	"runtime"
)

func execute() {
	//install := false

	flag.SetName("undo")
	flag.SetDescription("undo commands")
	flag.SetVersion("1.0.0")

	install := flag.NewSubcommand("install")
	install.Description = "Install undo command"
	uninstall := flag.NewSubcommand("uninstall")
	uninstall.Description = "Uninstall undo command"

	flag.AttachSubcommand(install, 1)
	flag.AttachSubcommand(uninstall, 1)
	flag.Parse()

	if install.Used {
		util.Install()
	} else if uninstall.Used {
		util.Uninstall()
	} else {
		util.Warn("undo is on development")
		util.Warn("It cannot undo all commands")
	}
}

func main() {
	if runtime.GOOS == "windows" {
		util.Error("Windows not supported")
	} else {
		execute()
	}
}
