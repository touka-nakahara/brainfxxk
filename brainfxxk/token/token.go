package token

type Tokentype byte

const (
	IncreasePointer Tokentype = iota
	DecreasePointer
	IncreaseValue
	DecreaseValue
	Output
	Input
	LoopForward
	LoopBackward

	Comment
)

var StringToTokenMap = map[byte]Tokentype{
	'>': IncreasePointer,
	'<': DecreasePointer,
	'+': IncreaseValue,
	'-': DecreaseValue,
	'.': Output,
	',': Input,
	'[': LoopForward,
	']': LoopBackward,
}

type Token struct {
	Tokentype Tokentype
	Literal   byte
}
