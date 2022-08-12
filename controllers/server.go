package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexFunc)
	router.HandleFunc("/api/v1/signup", signupfunc).Methods("POST") // TODO
	// router.HandleFunc("/api/v1/signin", signinFunc).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 第一引数にはHTTPサーバーからのレスポンスを出力することが出来るメソッドを持っている(該当のメソッドを実装している)構造体の値が来る
func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%T\n", w)                   // *http.response構造体
	fmt.Fprintf(w, "This is Go's Rest API") // メソッド内でw.Write()をするため
}

// 新規登録に必要な情報
type signupVaridation struct {
	FirstName string `validate:"required"`
	LastName  string `varidate:"required"`
	Email     string `varidate:"required,email"`
	Password  string `varidate:"required,lowercase,numeric,min=8,max=15"`
}

func signupfunc(w http.ResponseWriter, r *http.Request) {

	user := signupVaridation{
		FirstName: r.PostFormValue("firstname"),
		LastName:  r.PostFormValue("lastname"),
		Email:     r.PostFormValue("email"),
		Password:  r.PostFormValue("password"),
	}

	// TODO
	validate := validator.New()
	err := validate.Struct(user) // バリデーションの実行
	var output []string          // stringのスライス ["a","b","c"]

}
