package util

import "os"

func Exist(cmd *Command) bool {
	path := GetGoPath() + "/bin/undo-" + cmd.Main
	if _, err := os.Stat(path); err == nil {
		return true
	} else if !os.IsNotExist(err) {
		Error(true, err.Error())
	}
	return false
}
