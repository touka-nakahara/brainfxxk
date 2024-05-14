package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		// TODO: Add test cases.
		{
			name: "Hello World",
			args: []string{"cmd", "./input/mandelbrot.b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			main()
		})
	}
}

func Benchmark_main(b *testing.B) {
	tests := []struct {
		name string
		args []string
	}{
		// TODO: Add test cases.
		{
			name: "Hello World",
			args: []string{"cmd", "./input/prime.bf"},
		},
	}
	for _, tt := range tests {
		os.Args = tt.args
		originalStdOut := os.Stdout
		os.Stdout = nil
		defer func() {
			os.Stdout = originalStdOut
		}()
		main()
	}
}
