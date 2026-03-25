package lexer

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/mmnessim/go-stack/token"
	"github.com/mmnessim/go-stack/value"
)

type Lexer struct {
	Input    string
	Position int
}

func New(input string) *Lexer {
	return &Lexer{Input: input, Position: 0}
}

func (l *Lexer) Tokenize() ([]token.Token, error) {
	toks := []token.Token{}
	for l.Position < len(l.Input) {
		next, err := l.NextToken()
		if err != nil {
			return toks, err
		}
		toks = append(toks, next)
	}
	return toks, nil
}

func (l *Lexer) NextToken() (token.Token, error) {
	l.skipWhitespace()

	if l.Position >= len(l.Input) {
		return token.EOFToken(), nil
	}
	cur := l.Input[l.Position]
	switch cur {
	case '"':
		return l.readString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return l.readNumber()
	case '-':
		if l.Position+1 < len(l.Input) && unicode.IsDigit(rune(l.Input[l.Position+1])) {
			return l.readNumber()
		}
		return l.readWord()
	// TODO: add l.readComment() comments between ( and )
	default:
		return l.readWord()
	}
}

func (l *Lexer) skipWhitespace() {
	for l.Position < len(l.Input) && isWhitespace(l.Input[l.Position]) {
		l.Position += 1
	}
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\r' || char == '\n'
}

func (l *Lexer) readNumber() (token.Token, error) {
	start := l.Position
	if l.Input[l.Position] == '-' {
		l.Position++
	}
	for l.Position < len(l.Input) && unicode.IsDigit(rune(l.Input[l.Position])) {
		l.Position += 1
	}
	literal := l.Input[start:l.Position]
	num, err := strconv.ParseInt(literal, 10, 64)
	if err != nil {
		return token.ValueToken(nil), err
	}
	return token.ValueToken(value.Number{V: num}), nil
}

func (l *Lexer) readString() (token.Token, error) {
	l.Position++ // Opening "
	start := l.Position
	for l.Position < len(l.Input) && l.Input[l.Position] != '"' {
		l.Position++
	}
	if l.Position >= len(l.Input) {
		return token.ValueToken(nil), errors.New("unterminated string")
	}
	l.Position++ // Closing "
	return token.ValueToken(value.Str{V: l.Input[start : l.Position-1]}), nil
}

func (l *Lexer) readWord() (token.Token, error) {
	start := l.Position
	for l.Position < len(l.Input) && !unicode.IsSpace(rune(l.Input[l.Position])) {
		l.Position++
	}
	return token.WordToken(l.Input[start:l.Position]), nil
}
