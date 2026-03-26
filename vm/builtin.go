package vm

import (
	"errors"

	"github.com/mmnessim/go-stack/value"
)

func opAdd(vm *VM) error {
	x, y, err := vm.popTwoNumbers()
	if err != nil {
		return err
	}
	err = vm.Stack.Push(value.Number{V: x + y})
	if err != nil {
		return err
	}
	return nil
}

var typeError = errors.New("Type error")
