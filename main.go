package main

import (
	"fmt"
	interpreter "interpteter/interpriter"
	"io"
	"os"
	"time"
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
		fmt.Printf("os.Open: %s\n", err)
		return
	}

	buf, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("io.ReadAll: %s\n", err)
		return
	}

	// fmt.Printf("%s", string(buf))
	start := time.Now()
	// interpreter.InterpretProgram(string(buf))
	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s\n", elapsed)
	out := interpreter.Token(string(buf))

	fmt.Println(out)
	
	parser := interpreter.NewParser(out)
	res := parser.Parse()
	
	fmt.Println(res)

}
