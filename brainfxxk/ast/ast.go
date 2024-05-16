package ast

type Instruction interface {
	String() string
}

type IncreasePointerInstruction struct{}

func (l *IncreasePointerInstruction) String() string {
	return ">"
}

type DecreasePointerInstruction struct{}

func (l *DecreasePointerInstruction) String() string {
	return "<"
}

type IncreaseValueInstruction struct{}

func (l *IncreaseValueInstruction) String() string {
	return "+"
}

type DecreaseValueInstruction struct{}

func (l *DecreaseValueInstruction) String() string {
	return "-"
}

type OutputInstruction struct{}

func (l *OutputInstruction) String() string {
	return "."
}

type InputInstruction struct{}

func (l *InputInstruction) String() string {
	return ","
}

type LoopInstruction struct {
	Body []Instruction
}

func (l *LoopInstruction) String() string {
	output := "["
	for _, instruction := range l.Body {
		output += instruction.String()
	}
	output += "]"
	return output
}
