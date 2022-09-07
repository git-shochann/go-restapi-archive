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
	router.HandleFunc("/api/v1/signup", SignupFunc).Methods("POST")
	router.HandleFunc("/api/v1/signin", SigninFunc).Methods("POST")
	router.HandleFunc("/api/v1/create", CreateHabitFunc).Methods("POST")
	router.HandleFunc("/api/v1/update", UpdateHabitFunc).Methods("UPDATE")
	router.HandleFunc("/api/v1/delete/{id}", DeteteHabitFunc).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 第一引数にはHTTPサーバーからのレスポンスを出力することが出来るメソッドを持っている(該当のメソッドを実装している)構造体の値が来る
func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body: %v\n", r.Body)
	fmt.Printf("%T\n", w)                   // *http.response構造体
	fmt.Fprintf(w, "This is Go's Rest API") // メソッド内でw.Write()をするため
}
