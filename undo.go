package main

import (
	"github.com/buttercrab/undo/util"
	flag "github.com/integrii/flaggy"
	"os/exec"
	"runtime"
	"strings"
)

func execute() {
	flag.SetName("undo")
	flag.SetDescription("undo commands")
	flag.SetVersion("1.0.0")

	install := flag.NewSubcommand("install")
	install.Description = "Install undo command"
	uninstall := flag.NewSubcommand("uninstall")
	uninstall.Description = "Uninstall undo command"

	historyFile := ""
	flag.String(&historyFile, "", "history", "set shell's history file (unset if value is 'default')")
	startupFile := ""
	flag.String(&startupFile, "s", "start-up", "set shell's start-up file (unset if value is 'default')")

	flag.AttachSubcommand(install, 1)
	flag.AttachSubcommand(uninstall, 1)
	flag.Parse()

	if install.Used {
		util.Install()
	} else if uninstall.Used {
		util.Uninstall()
	} else {
		util.Warn("undo is under development")
		util.Warn("It cannot undo all commands")

		cmd := util.GetHistory(1)
		if !util.Exist(cmd) {
			util.Error(true, "Oops! Cannot undo")
		}

		output, err := exec.Command("undo-"+cmd.Main, "-m", `"`+cmd.Raw+`"`).Output()
		if err != nil {
			util.Error(true, err.Error())
		}

		for _, v := range strings.Split(string(output), "\n") {
			util.Log(v)
		}
	}
}

func main() {
	if runtime.GOOS == "windows" {
		util.Error(true, "Windows not supported")
	} else {
		execute()
	}
}
