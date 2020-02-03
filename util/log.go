package util

import (
	"fmt"
	"github.com/fatih/color"
	"math"
	"os"
)

const (
	escape = "\x1b"
)

var (
	block = []string{" ", "\u258F", "\u258E", "\u258D", "\u258C", "\u258B", "\u258A", "\u2589"}
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

func Check(err error) {
	if err != nil {
		Error(err.Error())
	}
}

func Log(s string) {
	prefix := color.CyanString("[LOG]")
	fmt.Printf("%s %s\n", prefix, s)
}

func printBar(progress float32, width int) {
	width *= 8
	count := int(math.Floor(float64(float32(width) * progress)))
	buf := make([]byte, 0, 100)
	for i := 0; i < width/8; i++ {
		if i < count/8 {
			buf = append(buf, "\u2588"...)
		} else if i == count/8 {
			buf = append(buf, block[count%8]...)
		} else {
			buf = append(buf, ' ')
		}
	}
	fmt.Printf("%s %.2f%%", string(buf), progress*100)
}

func PrintProgress() chan<- float32 {
	ch := make(chan float32, 1)
	go printProgress(ch)
	return ch
}

func printProgress(progress <-chan float32) {
	printBar(0, 72)
	for {
		cur := <-progress
		fmt.Print(escape + "[1K\r")
		if cur < 0 {
			printBar(1, 72)
			break
		} else {
			printBar(cur, 72)
		}
	}
	fmt.Println()
}
