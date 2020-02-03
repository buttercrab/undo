package util

import (
	"os"
	"os/exec"
)

func Install() {
	if os.Getenv("UNDO_INIT") == "true" {
		Warn("undo command already inited")
		return
	}

	file, err := os.OpenFile(GetHomePath()+"/."+GetDefaultShell()+"rc", os.O_APPEND|os.O_WRONLY, os.FileMode(0644))
	Check(err)
	defer file.Close()

	_, err = file.WriteString("\nexport UNDO_INIT=true;")
	Check(err)

	cmd := exec.Command("export", "UNDO_INIT=true")
	_, err = cmd.Output()
	Check(err)
}
