package ast

type Program struct {
	Source []Instruction
}

func (p *Program) String() string {
	output := ""
	for _, s := range p.Source {
		output += s.String()
	}
	return output
}
