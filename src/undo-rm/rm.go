package main

import (
	"github.com/buttercrab/undo/util"
	flag "github.com/integrii/flaggy"
	"os"
)

func Install() {
	err := os.Mkdir(util.GetHomePath()+"/.undo/data/rm/", os.FileMode(0644))
	if err != nil {
		util.Error(true, err.Error())
	}

	util.AppendToRC(`
function rm {
	$GOPATH/bin/rm do "$@"
}`)
}

func main() {
	do := flag.NewSubcommand("do")
	install := flag.NewSubcommand("install")
	undo := flag.NewSubcommand("undo")

	flag.AttachSubcommand(do, 1)
	flag.AttachSubcommand(install, 1)
	flag.AttachSubcommand(undo, 1)
	flag.Parse()

	if do.Used {

	} else if install.Used {
		Install()
	} else if undo.Used {

	} else {

	}
}
