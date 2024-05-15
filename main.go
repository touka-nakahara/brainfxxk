package main

import (
	brainfxxk "brainfxxk/interpreter"
	"fmt"
	"io"
	"os"
)

func main() {
	// 入力受け取り
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: arg - [filename] ")
		return
	}

	//TODO try catchに変更
	f, err := os.Open(args[1])

	if err != nil {
		fmt.Printf("os.Open: %s", err)
		return
	}

	buf, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("io.ReadAll: %s", err)
		return
	}

	fmt.Printf("%s", string(buf))

	brainfxxk.Interpreter(string(buf))
}
