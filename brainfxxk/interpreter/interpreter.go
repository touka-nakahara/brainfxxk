package brainfxxk

import (
	"context"
	"fmt"
	"time"
)

func Run(command string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resultCh := make(chan string)

	go func() {
		output := Interpreter(command)
		resultCh <- output
	}()

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("Interpreter Time Out")
	case output := <-resultCh:
		return output, nil
	}
}

func Interpreter(command string) string {
	var result string = ""
	var memory [10000000]int
	// in := bufio.NewReader(os.Stdin)
	pointer := 0
	idx := 0
	for idx < len(command) {
		switch command[idx] {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			// fmt.Printf("%c", memory[pointer])
			result += string(memory[pointer])
		case ',':
			// httpの場合は無視
			// byte, _ := in.ReadByte()
			// memory[pointer] = int(byte)
		case '[':
			if memory[pointer] == 0 {
				// 対応する]が見つかるまでidxを進める
				bracket_count := 1
				for {
					idx++
					if command[idx] == '[' {
						bracket_count++
					}
					if command[idx] == ']' {
						bracket_count--
					}
					if bracket_count == 0 {
						break
					}
				}
			}
		case ']':
			if memory[pointer] != 0 {
				//　対応する[が見つかるまで戻る
				bracket_count := 1
				for {
					idx--
					if command[idx] == '[' {
						bracket_count--
					}
					if command[idx] == ']' {
						bracket_count++
					}
					if bracket_count == 0 {
						break
					}
				}
			}
		}
		idx++
	}
	return result
}
