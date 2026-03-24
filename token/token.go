package token

import "github.com/mmnessim/go-stack/value"

type TokenType int

const (
	TokenValue TokenType = iota
	TokenWord
	TokenEOF
)

type Token struct {
	Type TokenType
	Val  value.Value
	Word string
}

func WordToken(w string) Token       { return Token{Type: TokenWord, Word: w} }
func ValueToken(v value.Value) Token { return Token{Type: TokenValue, Val: v} }
func EOFToken() Token                { return Token{Type: TokenEOF} }
