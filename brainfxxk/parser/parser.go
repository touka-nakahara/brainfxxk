package parser

import (
	"brainfxxk/brainfxxk/ast"
	"brainfxxk/brainfxxk/lexer"
	"brainfxxk/brainfxxk/token"
	"errors"
	"fmt"
	"io"
)

type Parser struct {
	lexer *lexer.Lexer
}

func NewParser(reader io.Reader) *Parser {
	l := lexer.NewLexer(reader)
	p := &Parser{lexer: l}
	return p
}

func (p *Parser) Parse() (*ast.Program, error) {
	var instructions = []ast.Instruction{}
	var stack = [][]ast.Instruction{instructions}

	for {
		tok, err := p.lexer.NextToken()
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
		// EOFの場合break
		if tok == nil {
			break
		}
		switch tok.Tokentype {
		case token.IncreasePointer:
			instructions = append(instructions, &ast.IncreasePointerInstruction{})
		case token.DecreasePointer:
			instructions = append(instructions, &ast.DecreasePointerInstruction{})
		case token.IncreaseValue:
			instructions = append(instructions, &ast.IncreaseValueInstruction{})
		case token.DecreaseValue:
			instructions = append(instructions, &ast.DecreaseValueInstruction{})
		case token.Output:
			instructions = append(instructions, &ast.OutputInstruction{})
		case token.Input:
			instructions = append(instructions, &ast.InputInstruction{})
		case token.LoopForward:
			loop := &ast.LoopInstruction{
				Body: []ast.Instruction{},
			}
			instructions = append(instructions, loop)
			stack = append(stack, instructions)
			instructions = loop.Body // Newするのめんどいから
		case token.LoopBackward:
			body := instructions
			if len(stack) == 1 {
				return nil, fmt.Errorf("unexpected Token: %c", tok.Literal)
			}
			instructions = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			lp := instructions[len(instructions)-1].(*ast.LoopInstruction)
			lp.Body = body
		}
	}
	if len(stack) != 1 {
		return nil, fmt.Errorf("expected Token: %c before end program", ']')
	}
	program := &ast.Program{
		Source: instructions,
	}
	return program, nil
}
