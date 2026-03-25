package lexer

import (
	"testing"

	"github.com/mmnessim/go-stack/token"
	"github.com/mmnessim/go-stack/value"
)

func TestNextToken_Number(t *testing.T) {
	l := &Lexer{Input: "42"}
	tok, err := l.NextToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.Type != token.TokenValue {
		t.Fatalf("expected TokenValue, got %v", tok.Type)
	}
	num, ok := tok.Val.(value.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", tok.Val)
	}
	if num.V != 42 {
		t.Errorf("expected 42, got %d", num.V)
	}
}

func TestNextToken_NegativeNumber(t *testing.T) {
	l := &Lexer{Input: "-42"}
	tok, err := l.NextToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := tok.Val.(value.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", tok.Val)
	}
	if num.V != -42 {
		t.Errorf("expected -42, got %d", num.V)
	}
}

func TestNextToken_NegativeSign_IsWord(t *testing.T) {
	l := &Lexer{Input: "- 42"}
	tok, err := l.NextToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.Type != token.TokenWord {
		t.Fatalf("expected TokenWord, got %v", tok.Type)
	}
	if tok.Word != "-" {
		t.Errorf("expected word '-', got %q", tok.Word)
	}
}

func TestNextToken_String(t *testing.T) {
	l := &Lexer{Input: `"hello"`}
	tok, err := l.NextToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	str, ok := tok.Val.(value.Str)
	if !ok {
		t.Fatalf("expected Str, got %T", tok.Val)
	}
	if str.V != "hello" {
		t.Errorf("expected 'hello', got %q", str.V)
	}
}

func TestNextToken_UnterminatedString(t *testing.T) {
	l := &Lexer{Input: `"hello`}
	_, err := l.NextToken()
	if err == nil {
		t.Fatal("expected error for unterminated string, got nil")
	}
}

func TestNextToken_Word(t *testing.T) {
	l := &Lexer{Input: "dup"}
	tok, err := l.NextToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.Type != token.TokenWord {
		t.Fatalf("expected TokenWord, got %v", tok.Type)
	}
	if tok.Word != "dup" {
		t.Errorf("expected 'dup', got %q", tok.Word)
	}
}

func TestNextToken_EOF(t *testing.T) {
	l := &Lexer{Input: ""}
	tok, err := l.NextToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.Type != token.TokenEOF {
		t.Errorf("expected TokenEOF, got %v", tok.Type)
	}
}

func TestTokenize_Expression(t *testing.T) {
	l := &Lexer{Input: "1 2 +"}
	toks, err := l.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(toks) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(toks))
	}

	cases := []struct {
		wantType token.TokenType
		wantVal  any
	}{
		{token.TokenValue, int64(1)},
		{token.TokenValue, int64(2)},
		{token.TokenWord, "+"},
	}

	for i, c := range cases {
		tok := toks[i]
		if tok.Type != c.wantType {
			t.Errorf("token %d: expected type %v, got %v", i, c.wantType, tok.Type)
		}
		switch want := c.wantVal.(type) {
		case int64:
			num, ok := tok.Val.(value.Number)
			if !ok {
				t.Errorf("token %d: expected Number, got %T", i, tok.Val)
			} else if num.V != want {
				t.Errorf("token %d: expected %d, got %d", i, want, num.V)
			}
		case string:
			if tok.Word != want {
				t.Errorf("token %d: expected word %q, got %q", i, want, tok.Word)
			}
		}
	}
}

func TestTokenize_SkipsWhitespace(t *testing.T) {
	l := &Lexer{Input: "  42   dup  "}
	toks, err := l.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 42 dup EOF
	if len(toks) != 3 {
		t.Fatalf("expected 2 tokens, got %d", len(toks))
	}
}

func TestTokenize_StringWithSpaces(t *testing.T) {
	l := &Lexer{Input: `"hello world"`}
	toks, err := l.Tokenize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(toks) != 1 {
		t.Fatalf("expected 1 token, got %d", len(toks))
	}
	str, ok := toks[0].Val.(value.Str)
	if !ok {
		t.Fatalf("expected Str, got %T", toks[0].Val)
	}
	if str.V != "hello world" {
		t.Errorf("expected 'hello world', got %q", str.V)
	}
}
