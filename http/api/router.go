package api

import (
	"net/http"
)

func NewRouter() http.Handler {
	m := http.NewServeMux()

	// ハンドラ設定
	m.Handle("GET /", http.FileServer(http.Dir("http/static/root")))
	m.HandleFunc("POST /run", BrainfxxkFuncPOST)

	// タイムアウト ( 2分 )
	//TODO 503が返っちゃう
	// h := http.TimeoutHandler(m, 2*time.Minute, "Request timed out.")

	return m
}
