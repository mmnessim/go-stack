package repl

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mmnessim/go-stack/stack"
	"github.com/mmnessim/go-stack/value"
)

func Repl() {
	s := stack.New()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "bye" {
			break
		}
		if n, err := strconv.ParseInt(line, 10, 64); err == nil {
			s.Push(value.Number{V: n})
		} else {
			continue
		}
		s.Print()
	}
}
