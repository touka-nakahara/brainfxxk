package lexer

import (
	"io"

	"brainfxxk/brainfxxk/token"
)

type Lexer struct {
	reader io.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	l := &Lexer{
		reader: reader,
	}
	return l
}

func (l *Lexer) NextToken() (*token.Token, error) {
	b := make([]byte, 1)
	_, err := l.reader.Read(b)
	if err != nil {
		return nil, err
	}

	tokenType := token.Comment

	if token, ok := token.StringToTokenMap[b[0]]; ok {
		tokenType = token
	}

	token := &token.Token{
		Tokentype: tokenType,
		Literal:   b[0],
	}

	return token, nil
}
