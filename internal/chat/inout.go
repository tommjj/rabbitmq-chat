package chat

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return ""
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return line
}
