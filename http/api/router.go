package api

import (
	"net/http"
)

func NewRouter() http.Handler {
	m := http.NewServeMux()

	// ハンドラ設定
	m.Handle("GET /", http.FileServer(http.Dir("http/static/root")))
	m.HandleFunc("POST /run", BrainfxxkFuncPOST)

	return m
}
