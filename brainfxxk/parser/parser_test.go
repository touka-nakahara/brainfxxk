package parser

import (
	"brainfxxk/brainfxxk/ast"
	"fmt"
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		input    string
		expected []ast.Instruction
	}{
		{
			input:    "...[++>],.+[]",
			expected: []ast.Instruction{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			parser := NewParser(input)
			program, err := parser.Parse()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(program)
			}
		})
	}
}
