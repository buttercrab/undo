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
		Error("We don't support " + shell + " yet.")
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
