package util

import "os"

func GetHistory(before int) *Command {
	shell := GetDefaultShell()

	var s string
	switch shell {
	case "zsh":
		s = zshHistory(before)
	case "bash":
		s = bashHistory(before)
	default:
		Error(true, "We don't support %s yet.", shell)
	}

	return NewCommand(s)
}

func zshHistory(before int) string {
	file, err := os.Open(GetHistFile())
	Check(err)
	defer file.Close()

	// TODO

	return ""
}

func bashHistory(before int) string {
	file, err := os.Open(GetHistFile())
	Check(err)
	defer file.Close()

	// TODO

	return ""
}
