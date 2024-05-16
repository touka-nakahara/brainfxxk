package server

import "time"

type ServerConfig struct {
	Address     string
	IdleTimeout time.Duration
}
