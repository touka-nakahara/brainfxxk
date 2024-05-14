package parser

import "go/token"

const (
	op_inc_pt  = "IncPointer"
	op_dec_pt  = "DecPointer"
	op_inc_val = "IncreValue"
	op_dec_val = "DecValue"
	op_out     = "OutputValue"
	op_in      = "InputValue"
	op_jmp_bck = "LoopStart"
	op_jmp_fwd = "LoopEnd"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

type TokenType string

type Token struct {
	Type TokenType
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

type Node interface {
	TokenLiteral() string
}

type Instrcution interface {
	Node
}

type Program struct {
	instrcution []Instrcution
}

func (p *Program) TokenLiteral() string {
	if len(p.instrcution) > 0 {
		return p.instrcution[0].TokenLiteral()
	} else {
		return ""
	}
}

type AddPointer struct {
	Token token.Token
}
