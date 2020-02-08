package main

import (
	"github.com/buttercrab/undo/util"
	flag "github.com/integrii/flaggy"
)

func install() {
	util.AppendToRC(`
function rm {
	$GOPATH/bin/rm do "$@"
}`)
}

func main() {
	do := flag.NewSubcommand("do")
	undo := flag.NewSubcommand("undo")

	flag.AttachSubcommand(do, 1)
	flag.AttachSubcommand(undo, 1)
	flag.Parse()

	if do.Used {

	} else if undo.Used {

	} else {

	}
}
