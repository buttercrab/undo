package util

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Warn(s string) {
	prefix := color.YellowString("[WARN]")
	fmt.Printf("%s %s\n", prefix, s)
}

func Error(s string) {
	prefix := color.RedString("[ERROR]")
	fmt.Printf("%s %s\n", prefix, s)
	Warn("Exiting program by error")
	os.Exit(1)
}

func Log(s string) {
	prefix := color.CyanString("[LOG]")
	fmt.Printf("%s %s\n", prefix, s)
}
