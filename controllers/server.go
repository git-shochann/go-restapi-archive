package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexFunc)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 第一引数にはHTTPサーバーからのレスポンスを出力することが出来るメソッドを持っている構造体の値が来る
func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%T\n", w)                   // *http.response
	fmt.Fprintf(w, "This is Go's Rest API") // メソッド内でw.Write()をするため
}
