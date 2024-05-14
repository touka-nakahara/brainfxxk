package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const MEMORYMAX int = 10000

func InterpretProgram(command string) {
	var memory [MEMORYMAX]int
	in := bufio.NewReader(os.Stdin)
	pointer := 0
	program_counter := 0
	for program_counter < len(command) {
		switch command[program_counter] {
		case '>':
			pointer++
			// if pointer >= MEMORYMAX {
			// 	fmt.Printf("Error: Out of Range: Memory Access :[%d]\n", pointer)
			// 	return
			// }
		case '<':
			pointer--
			// if pointer < 0 {
			// 	fmt.Printf("Error: Out of Ranger Memory Access :[%d]\n", pointer)
			// 	return
			// }
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			fmt.Printf("%c", memory[pointer])
		case ',':
			byte, err := in.ReadByte()
			if err == io.EOF {
				return
			}
			memory[pointer] = int(byte)
		case '[':
			if memory[pointer] == 0 {
				// 対応する]が見つかるまでidxを進める
				bracket_count := 1
				for {
					program_counter++
					// if program_counter >= MEMORYMAX {
					// 	fmt.Printf("Error: Out of Range: Memory Access :[%d]\n", pointer)
					// 	return
					// }
					if command[program_counter] == '[' {
						bracket_count++
					}
					if command[program_counter] == ']' {
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
					program_counter--
					// if program_counter < 0 {
					// 	fmt.Printf("Error: Out of Ranger Memory Access :[%d]\n", pointer)
					// 	return
					// }
					if command[program_counter] == '[' {
						bracket_count--
					}
					if command[program_counter] == ']' {
						bracket_count++
					}
					if bracket_count == 0 {
						break
					}
				}
			}
		}
		program_counter++
	}
	fmt.Println("")
}
