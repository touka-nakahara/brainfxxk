package api

import (
	"fmt"
	"io"
	"net/http"

	brainfxxk "brainfxxk/brainfxxk/interpreter"
)

func BrainfxxkFuncPOST(w http.ResponseWriter, r *http.Request) {
	// bodyの解釈
	//TODO ここで全部読み込んだ方がいい気がする (入力の読み込みのエラーとBrainFxxkの読み込みを分けたい)
	//TODO r.Bodyがめっちゃデカかったら壊れそう
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	command := string(body)

	// brainfxxkの起動
	result, err := brainfxxk.Run(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// レスポンス返答(text/plain)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintln(w, result)
}
