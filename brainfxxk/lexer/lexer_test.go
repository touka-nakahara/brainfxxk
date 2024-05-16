package lexer

import (
	"brainfxxk/brainfxxk/token"
	"io"
	"strings"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	type fields struct {
		input io.Reader
	}
	tests := []struct {
		input    string
		expected []*token.Token
	}{
		{
			input: "><+-,.[]",
			expected: []*token.Token{
				{Tokentype: token.IncreasePointer, Literal: '>'},
				{Tokentype: token.DecreasePointer, Literal: '<'},
				{Tokentype: token.IncreaseValue, Literal: '+'},
				{Tokentype: token.DecreaseValue, Literal: '-'},
				{Tokentype: token.Input, Literal: ','},
				{Tokentype: token.Output, Literal: '.'},
				{Tokentype: token.LoopForward, Literal: '['},
				{Tokentype: token.LoopBackward, Literal: ']'},
			},
		},
		{
			input: "@:;{}",
			expected: []*token.Token{
				{Tokentype: token.Comment, Literal: '@'},
				{Tokentype: token.Comment, Literal: ':'},
				{Tokentype: token.Comment, Literal: ';'},
				{Tokentype: token.Comment, Literal: '{'},
				{Tokentype: token.Comment, Literal: '}'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			l := &Lexer{
				reader: input,
			}
			for _, expected := range tt.expected {
				token, err := l.NextToken()
				if err != nil {
					t.Fatalf("unexpected Error %v", err)
				}
				if token.Tokentype != expected.Tokentype {
					t.Errorf("exptectet: %v, but got: %v", expected.Tokentype, token.Tokentype)
				}
				if token.Literal != expected.Literal {
					t.Errorf("exptectet: %v, but got: %v", expected.Literal, token.Literal)
				}
			}
		})
	}
}
