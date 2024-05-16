package server

import (
	"context"
	"errors"
	"fmt"
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
		// 通常のシャットダウンチェック
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()                                                            // キャンセル通知が来たら
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // タイムアウトのコンテキスト生成s
	defer cancel()
	s.Shutdown(ctx) // アイドル状態になってからシャットダウン
	fmt.Println("Server: graceful shutdown")

}
