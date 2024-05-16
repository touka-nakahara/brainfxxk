package expr

import (
	"brainfxxk/brainfxxk/ast"
	"brainfxxk/brainfxxk/parser"
	"io"
)

func Run(input io.Reader) (string, error) {
	// パーサー作って
	p := parser.NewParser(input)
	program, err := p.Parse()
	// 構文エラー
	if err != nil {
		return "", err
	}
	// 実行
	expr := NewExpr()
	output, err := expr.Run(program)
	// 実行時エラー
	if err != nil {
		return "", err
	}

	// 値を返す (string)
	return output, nil
}

type Expr struct{}

func NewExpr() *Expr {
	e := &Expr{}
	return e
}

func (e *Expr) Run(program *ast.Program) (string, error) {
	return "", nil
}
