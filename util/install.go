package util

import (
	"os"
)

func Install() {
	if IsInstalled() {
		Error(true, "undo had already inited")
	}

	Log("checking your shell: %s", GetDefaultShell())
	Log("installing undo")

	ch := PrintProgress()

	ch <- -1

	Info("Please restart your terminal")
}

func AppendToRC(s string) {
	file, err := os.OpenFile(GetStartFile(), os.O_APPEND|os.O_WRONLY, os.FileMode(0644))
	Check(err)
	defer file.Close()

	_, err = file.WriteString("\n" + s)
	Check(err)
}
