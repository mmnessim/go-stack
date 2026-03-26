package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mmnessim/go-stack/lexer"
	"github.com/mmnessim/go-stack/stack"
	"github.com/mmnessim/go-stack/vm"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	l := lexer.New("")
	vm := vm.VM{Stack: *stack.New()}

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			vm.Stack.Print()
			continue
		}
		if line == "bye" {
			break
		}

		l.Input = line
		l.Position = 0
		toks, err := l.Tokenize()
		if err != nil {
			fmt.Println(err)
			continue
		}
		vm.Eval(toks)

		vm.Stack.Print()
	}
}
