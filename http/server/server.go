package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"brainfxxk/http/api"
)

func StartServer() {
	// interruptシグナルを受信したときに、コンテキストにキャンセルを通知する
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	r := api.NewRouter()

	s := http.Server{
		Addr:        ":80",
		IdleTimeout: 24 * time.Hour,
		Handler:     r,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.Shutdown(ctx)
	}()

	// 通常のシャットダウンチェック
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

}
