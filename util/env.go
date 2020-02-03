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
