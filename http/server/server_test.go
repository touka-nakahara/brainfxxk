package server

import "testing"

func TestServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "serve"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartServer()
		})
	}
}
