// a simple print with color

package print

import (
	"fmt"
	"io"
)

type Color string

const (
	Reset   Color = "\033[0m"
	Unset   Color = ""
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	Gray    Color = "\033[37m"
	White   Color = "\033[97m"
)

func (c Color) set() {
	if c == "" {
		return
	}

	fmt.Print(c)
}

func (c Color) reset() {
	if c == "" {
		return
	}

	fmt.Print(Reset)
}

func (c Color) Print(a ...any) {
	c.set()
	defer c.reset()

	fmt.Print(a...)
}

func (c Color) Printf(format string, a ...any) {
	c.set()
	defer c.reset()

	fmt.Printf(format, a...)
}

func (c Color) Println(a ...any) {
	c.set()
	defer c.reset()

	fmt.Println(a...)
}

type Print struct {
	Output io.Writer
	Color  Color
}

func (p *Print) reset() {
	if p.Color == "" {
		return
	}

	fmt.Fprint(p.Output, Reset)
}

func (p *Print) set() {
	if p.Color == "" {
		return
	}

	fmt.Fprint(p.Output, p.Color)
}

func (p *Print) Print(a ...any) {
	p.set()
	defer p.reset()

	fmt.Fprint(p.Output, a...)
}

func (p *Print) Printf(format string, a ...any) {
	p.set()
	defer p.reset()

	fmt.Fprintf(p.Output, format, a...)
}

func (p *Print) Println(a ...any) {
	p.set()
	defer p.reset()

	fmt.Fprintln(p.Output, a...)
}
