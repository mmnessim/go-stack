package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	if err := s.Push(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s.Top != 1 {
		t.Errorf("expected Top=1, got %d", s.Top)
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	val, err := s.Pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 2 {
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
	for i := range CAPACITY {
		s.Push(i)
	}
	if err := s.Push(999); err != ErrStackOverflow {
		t.Errorf("expected ErrStackOverflow, got %v", err)
	}
}

func TestSwap(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Swap()
	val, _ := s.Pop()
	if val != 1 {
		t.Errorf("expected 1 on top after swap, got %d", val)
	}
}

func TestDup(t *testing.T) {
	s := New()
	s.Push(1)
	s.Dup()
	if s.Top != 2 {
		t.Errorf("expected s.Top=2, got %d", s.Top)
	}
}
