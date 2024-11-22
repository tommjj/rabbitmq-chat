package pcolor

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {

	Red.Println("hello world")

	Blue.Println("hello world")

	fmt.Print(Yellow.Sprintln("hello world", ""))
}
