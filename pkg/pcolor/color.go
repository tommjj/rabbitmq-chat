// a simple print with color

package pcolor

import (
	"fmt"
)

type color string

const (
	Reset   color = "\033[0m"
	Unset   color = ""
	Red     color = "\033[31m"
	Green   color = "\033[32m"
	Yellow  color = "\033[33m"
	Blue    color = "\033[34m"
	Magenta color = "\033[35m"
	Cyan    color = "\033[36m"
	Gray    color = "\033[37m"
	White   color = "\033[97m"
)

func (c color) set() {
	if c == "" {
		return
	}

	fmt.Print(c)
}

func (c color) reset() {
	if c == "" {
		return
	}

	fmt.Print(Reset)
}

func (c color) Print(a ...any) (n int, err error) {
	c.set()
	defer c.reset()

	return fmt.Print(a...)
}

func (c color) Printf(format string, a ...any) (n int, err error) {
	c.set()
	defer c.reset()

	return fmt.Printf(format, a...)
}

func (c color) Println(a ...any) (n int, err error) {
	c.set()
	defer c.reset()

	return fmt.Println(a...)
}

func (c color) Sprint(a ...any) string {
	if c == "" {
		return fmt.Sprint(a...)
	}

	return fmt.Sprint(c, fmt.Sprint(a...), Reset)
}

func (c color) Sprintf(format string, a ...any) string {
	if c == "" {
		return fmt.Sprintf(format, a...)
	}

	return fmt.Sprint(c, fmt.Sprintf(format, a...), Reset)
}

func (c color) Sprintln(a ...any) string {
	if c == "" {
		return fmt.Sprintln(a...)
	}

	return fmt.Sprint(c, fmt.Sprint(a...), Reset, "\n")
}
