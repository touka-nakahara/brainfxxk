package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	brainfxxk "brainfxxk/brainfxxk/interpreter"
)

func IndexHandleFuncGET(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "http/static/root/")
	}
}

func BrainfxxkFuncPOST(w http.ResponseWriter, r *http.Request) {
	// bodyの解釈
	//TODO ここで全部読み込んだ方がいい気がする (入力の読み込みのエラーとBrainFxxkの読み込みを分けたい)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		//TODO Request Errorを返したい
	}
	command := string(body)

	// brainfxxkの起動
	//TODO レスポンスとBFの起動は分けたい
	result, err := brainfxxk.Run(command)
	if err != nil {
		//TODO 403を返す
	}

	// レスポンス返答(text/plain)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintln(w, result)
}

func Server() {

	s := http.Server{}

	//TODO forループでhanderスライスを見て, ハンドル登録する
	http.Handle("GET /", http.FileServer(http.Dir("http/static/root")))
	http.HandleFunc("POST /run", BrainfxxkFuncPOST)
	// Methodが読めるらしい

	err := s.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
