package util

import (
	"net/http"
	"os"
)

func Install() {
	if IsInstalled() {
		Error(true, "undo had already inited")
	}

	Log("checking your shell: %s", GetDefaultShell())
	Log("fetching command list")

	res, err := http.Get(GetBaseUrl() + "/blob/master/commands.txt")
	if err != nil {
		Error(true, err.Error())
	}
	defer res.Body.Close()

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
