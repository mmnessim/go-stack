package stack

import (
	"testing"

	"github.com/mmnessim/go-stack/value"
	v "github.com/mmnessim/go-stack/value"
)

func TestPush(t *testing.T) {
	s := New()
	if err := s.Push(value.Number{V: 1}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.Top != 1 {
		t.Errorf("expected Top=1, got %d", s.Top)
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(value.Number{V: 1})
	s.Push(value.Number{V: 2})
	val, err := s.Pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != (value.Number{V: 2}) {
		t.Errorf("expected 2, got %d", val)
	}
}

func TestStackUnderflow(t *testing.T) {
	s := New()
	_, err := s.Pop()
	if err != ErrStackUnderflow {
		t.Errorf("expected ErrStackUnderflow, got %v", err)
	}
}

func TestStackOverflow(t *testing.T) {
	s := New()
	for _ = range CAPACITY {
		s.Push(v.Number{V: 1})
	}
	if err := s.Push(v.Number{V: 999}); err != ErrStackOverflow {
		t.Errorf("expected ErrStackOverflow, got %v", err)
	}
}

func TestSwap(t *testing.T) {
	s := New()
	s.Push(value.Number{V: 1})
	s.Push(value.Number{V: 2})
	s.Swap()
	val, _ := s.Pop()
	if val != (value.Number{V: 1}) {
		t.Errorf("expected 1 on top after swap, got %d", val)
	}
}

func TestDup(t *testing.T) {
	s := New()
	s.Push(value.Number{V: 1})
	s.Dup()
	if s.Top != 2 {
		t.Errorf("expected s.Top=2, got %d", s.Top)
	}
}
