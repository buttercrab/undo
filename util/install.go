package util

import (
	"os"
	"os/exec"
	"strings"
)

func getDefaultShell() string {
	list := strings.Split(os.Getenv("SHELL"), "/")
	return list[len(list)-1]
}

func getHomePath() string {
	return os.Getenv("HOME")
}

func Install() {
	if os.Getenv("UNDO_INIT") == "true" {
		Warn("undo command already inited")
		return
	}

	file, err := os.OpenFile(getHomePath()+"/."+getDefaultShell()+"rc", os.O_APPEND|os.O_WRONLY, os.FileMode(0644))

	if err != nil {
		Error(err.Error())
	}

	defer file.Close()

	if _, err := file.WriteString("\nexport UNDO_INIT=true;"); err != nil {
		Error(err.Error())
	}

	cmd := exec.Command("export", "UNDO_INIT=true")
	_, err = cmd.Output()

	if err != nil {
		Error(err.Error())
	}
}
