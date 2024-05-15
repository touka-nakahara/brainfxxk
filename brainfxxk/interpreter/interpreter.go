package brainfxxk

import (
	"bufio"
	"fmt"
	"os"
)

func Interpreter(command string) {
	var memory [128]int
	in := bufio.NewReader(os.Stdin)
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
			fmt.Printf("%c", memory[pointer])
		case ',':
			byte, _ := in.ReadByte()
			memory[pointer] = int(byte)
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
	fmt.Println("")
}
