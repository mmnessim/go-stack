package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mmnessim/go-stack/lexer"
	"github.com/mmnessim/go-stack/vm"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	l := lexer.New("")
	v := vm.New()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			v.Stack.Print()
			continue
		}
		l.Input = line
		l.Position = 0

		toks, err := l.Tokenize()
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		if err := v.Eval(toks); err != nil {
			fmt.Println("error:", err)
			continue
		}

		v.Stack.Print()
	}
}
