package stack

import (
	"errors"
	"fmt"

	v "github.com/mmnessim/go-stack/value"
)

const CAPACITY = 256

type Stack struct {
	Items [CAPACITY]v.Value
	Top   uint
}

func New() *Stack {
	return &Stack{Top: 0}
}

func (s *Stack) Push(val v.Value) error {
	if s.Top >= CAPACITY {
		return ErrStackOverflow
	}
	s.Items[s.Top] = val
	s.Top += 1
	return nil
}

func (s *Stack) Pop() (v.Value, error) {
	if s.Top == 0 {
		return nil, ErrStackUnderflow
	}
	val := s.Items[s.Top-1]
	s.Top -= 1
	return val, nil
}

func (s *Stack) Peek() (v.Value, error) {
	if s.Top == 0 {
		return nil, ErrStackUnderflow
	}
	return s.Items[s.Top-1], nil
}

func (s *Stack) Swap() error {
	if s.Top < 2 {
		return ErrStackUnderflow
	}
	s.Items[s.Top-1], s.Items[s.Top-2] = s.Items[s.Top-2], s.Items[s.Top-1]
	return nil
}

func (s *Stack) Dup() error {
	if s.Top == 0 {
		return ErrStackUnderflow
	}
	s.Items[s.Top] = s.Items[s.Top-1]
	s.Top += 1
	return nil
}

func (s *Stack) Print() {
	if s.Top == 0 {
		return
	}
	fmt.Print("  ")
	for idx := range s.Top {
		fmt.Print(s.Items[idx].String(), " ")
	}
	fmt.Print("<- Top\n")
}

var ErrStackOverflow = errors.New("Stack overflow")
var ErrStackUnderflow = errors.New("Stack underflow")
