package vm

import (
	"fmt"
	"os"

	"github.com/mmnessim/go-stack/value"
)

func opPop(vm *VM) error {
	x, err := vm.Stack.Pop()
	if err != nil {
		return err
	}
	fmt.Println(x.String())
	return nil
}

func opBye(vm *VM) error {
	os.Exit(0)
	return nil
}

/// Math

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

func opSubtract(vm *VM) error {
	x, y, err := vm.popTwoNumbers()
	if err != nil {
		return err
	}
	err = vm.Stack.Push(value.Number{V: y - x})
	if err != nil {
		return err
	}
	return nil
}

func opMult(vm *VM) error {
	x, y, err := vm.popTwoNumbers()
	if err != nil {
		return err
	}
	err = vm.Stack.Push(value.Number{V: y * x})
	if err != nil {
		return err
	}
	return nil
}

func opDivide(vm *VM) error {
	x, y, err := vm.popTwoNumbers()
	if err != nil {
		return err
	}
	err = vm.Stack.Push(value.Number{V: y / x})
	if err != nil {
		return err
	}
	return nil
}
