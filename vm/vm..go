package vm

import (
	"github.com/mmnessim/go-stack/stack"
	"github.com/mmnessim/go-stack/token"
)

type VM struct {
	Stack stack.Stack
}

func (vm *VM) Eval(tokens []token.Token) error {
	for _, tok := range tokens {
		switch tok.Type {
		case token.TokenValue:
			{
				err := vm.Stack.Push(tok.Val)
				if err != nil {
					return err
				}
			}
		case token.TokenWord:
			{
				err := vm.execWord(tok.Word)
				if err != nil {
					return err
				}
			}
		case token.TokenEOF:
			{
				break
			}
		}

	}

	return nil
}

func (vm *VM) execWord(word string) error {
	return nil
}
