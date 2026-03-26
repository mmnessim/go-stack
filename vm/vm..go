package vm

import (
	"fmt"

	"github.com/mmnessim/go-stack/stack"
	"github.com/mmnessim/go-stack/token"
	"github.com/mmnessim/go-stack/value"
)

type VM struct {
	Stack      stack.Stack
	Dictionary map[string]func(*VM) error
}

func New() *VM {
	vm := &VM{
		Stack:      *stack.New(),
		Dictionary: make(map[string]func(*VM) error),
	}
	vm.registerBuiltins()
	return vm
}

func (vm *VM) registerBuiltins() {
	vm.Dictionary["+"] = opAdd
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
				err := vm.ExecWord(tok.Word)
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

func (vm *VM) ExecWord(word string) error {
	op, ok := vm.Dictionary[word]
	if !ok {
		return fmt.Errorf("unknown word: %s", word)
	}
	return op(vm)
}
func (vm *VM) popTwo() (value.Value, value.Value, error) {
	if vm.Stack.Top < 2 {
		return nil, nil, stack.ErrStackUnderflow
	}
	valOne, err := vm.Stack.Pop()
	if err != nil {
		return nil, nil, err
	}
	valTwo, err := vm.Stack.Pop()
	if err != nil {
		return nil, nil, err
	}
	return valOne, valTwo, nil
}

func (vm *VM) popTwoNumbers() (int64, int64, error) {
	x, y, err := vm.popTwo()
	if err != nil {
		return 0, 0, err
	}
	xn, ok1 := x.(value.Number)
	yn, ok2 := y.(value.Number)
	if !ok1 || !ok2 {
		return 0, 0, typeError
	}
	return xn.V, yn.V, nil
}

func (vm *VM) popTwoStrings() (string, string, error) {
	x, y, err := vm.popTwo()
	if err != nil {
		return "", "", err
	}
	xstr, ok1 := x.(value.Str)
	ystr, ok2 := y.(value.Str)
	if !ok1 || !ok2 {
		return "", "", err
	}
	return xstr.V, ystr.V, nil
}
