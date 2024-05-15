package server

import (
	"log"
	"net/http"
)

type indexHandler struct{}
type brainfxxkHadler struct{}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// GET Method のみ対応
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	//TODO どうしたらいいんやこれ...
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/root")
	} else {
		//TODO 最終ハンドラでNotFoundを見るのってほんとなんかな...
		http.NotFound(w, r)
	}
}

func (h *brainfxxkHadler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// POSTのみ対応
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	// bodyの解釈
	// brainfxxkの起動
	// レスポンス返答
	// エラーの場合は403
}

func Server() {

	s := http.Server{}

	//TODO forループでhanderスライスを見て, ハンドル登録する
	http.Handle("/", new(indexHandler))
	// Methodが読めるらしい

	err := s.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
