package util

import (
	"os"
	"strings"
)

func GetDefaultShell() string {
	list := strings.Split(os.Getenv("SHELL"), "/")
	return list[len(list)-1]
}

func GetHomePath() string {
	return os.Getenv("HOME")
}

func GetHistFile() string {
	return os.Getenv("HISTFILE")
}

func GetGoPath() string {
	return os.Getenv("GOPATH")
}

func GetStartFile() string {
	shell := GetDefaultShell()
	home := GetHomePath()

	res := ""
	var list []string
	switch shell {
	case "bash":
		list = []string{"/.bashrc", "/.bash_profile"}
		break
	case "zsh":
		list = []string{"/.zshrc", "/.zprofile"}
		break
	default:
		Error(true, "We don't support %s yet", shell)
	}

	for _, file := range list {
		if _, err := os.Stat(home + file); err == nil {
			res = home + file
			return res
		} else if !os.IsNotExist(err) {
			Error(true, err.Error())
		}
	}

	Error(false, "We couldn't find your shell start-up file")
	Error(true, "Type 'undo -s FILE_NAME' to set your start-up file")

	return res
}

func IsInstalled() bool {
	return true
}
