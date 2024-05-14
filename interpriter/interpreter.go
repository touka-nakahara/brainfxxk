package interpreter

const (
	op_inc_pt = iota
	op_dec_pt
	op_inc_val
	op_dec_val
	op_out
	op_in
	op_jmp_fwd
	op_jmp_bck
)

// 字句解析
func Token(command string) []rune {
	var instrcution []rune
	for _, cmd := range command {
		switch cmd {
		case '>':
			instrcution = append(instrcution, op_inc_pt)
		case '<':
			instrcution = append(instrcution, op_dec_pt)
		case '+':
			instrcution = append(instrcution, op_inc_val)
		case '-':
			instrcution = append(instrcution, op_dec_val)
		case '.':
			instrcution = append(instrcution, op_out)
		case ',':
			instrcution = append(instrcution, op_in)
		case '[':
			instrcution = append(instrcution, op_jmp_fwd)
		case ']':
			instrcution = append(instrcution, op_jmp_bck)
		}
	}
	return instrcution
}

type Instruction struct {
	instruction []interface{} // どうしようもない
}

func NewInstruction() *Instruction {
	i := &Instruction{}
	return i
}

type Parser struct {
	input    []rune // トークン化された命令
	position int    // 現在の読み込み場所
}

func NewParser(input []rune) *Parser {
	p := &Parser{input: input, position: 0}
	return p
}

func (p *Parser) Parse() *Instruction {
	mainInst := NewInstruction()
	stack := []*Instruction{mainInst}

	for p.position < len(p.input) {
		token := p.input[p.position]
		p.position++

		switch token {
		case op_jmp_fwd:
			nastedInst := NewInstruction()
			currentInst := stack[len(stack)-1]
			currentInst.instruction = append(currentInst.instruction, nastedInst)
			stack = append(stack, nastedInst)
		case op_jmp_bck:
			stack = stack[:len(stack)-1]
		default:
			currentInst := stack[len(stack)-1]
			currentInst.instruction = append(currentInst.instruction, token)
		}

	}
	return mainInst
}
