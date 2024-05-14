package interpreter

const (
	op_inc_pt = iota
	op_dec_pt
	op_inc_val
	op_dec_val
	op_out
	op_in
	op_jmp_bck
	op_jmp_fwd
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
	Instruction []interface{} // どうしようもない
}

type Parser struct {
	input    []rune // トークン化された命令
	position int    // 現在の読み込み場所
}

func NewParser(input []rune) *Parser {
	p := &Parser{input: input, position: 0}
	return p
}

// func New() *Instrcution {
// 	i := &Instrcution{}
// 	return i
// }

// func (p *Parser) Parse(inst Instrcution) Instrcution {
// 	for p.position < len(p.input) {
// 		token := p.input[p.position]
// 		switch token {
// 		case 6:
// 			p.Parse(*inst.childlen)
// 		case 7:
// 			return inst
// 		default:
// 			inst.instruction = append(inst.instruction, token)
// 		}
// 		p.position++
// 	}
// 	return inst
// }

// // func Parser(tokens []rune) {
// // 	instructions := New()
// // 	for _, token := range tokens {
// 		switch token {
// 		case 6:
// 			instructions.childlen
// 		case 7:
// 			return
// 		default:
// 			instructions.instruction = append(instructions.instruction, token)
// 		}
// 	}
// }
