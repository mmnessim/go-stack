package stack

import "errors"

const CAPACITY = 256

type Stack struct {
	Items [CAPACITY]int
	Top   uint
}

func New() *Stack {
	return &Stack{Top: 0}
}

func (s *Stack) Push(val int) error {
	if s.Top >= CAPACITY {
		return ErrStackOverflow
	}
	s.Items[s.Top] = val
	s.Top += 1
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Top == 0 {
		return 0, ErrStackUnderflow
	}
	val := s.Items[s.Top-1]
	s.Top -= 1
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if s.Top < 2 {
		return 0, ErrStackUnderflow
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

var ErrStackOverflow = errors.New("Stack overflow")
var ErrStackUnderflow = errors.New("Stack underflow")
