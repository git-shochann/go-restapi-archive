package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexFunc)
	router.HandleFunc("/api/v1/signup", signupFunc).Methods("POST")
	router.HandleFunc("/api/v1/signin", signinFunc).Methods("POST")
	router.HandleFunc("api/v1/create", createHabitFunc).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 第一引数にはHTTPサーバーからのレスポンスを出力することが出来るメソッドを持っている(該当のメソッドを実装している)構造体の値が来る
func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Body: %v\n", r.Body)
	fmt.Printf("%T\n", w)                   // *http.response構造体
	fmt.Fprintf(w, "This is Go's Rest API") // メソッド内でw.Write()をするため
}

func signupFunc(w http.ResponseWriter, r *http.Request) {

	// Jsonでくるので、まずGoで使用できるようにする
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var user models.UserSignupVaridation
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		models.SendResponse(w, "Unable to unmarshal json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	fmt.Printf("user: %v\n", user)

	ok, result := user.SignupVaridator()

	// false時の処理
	if !ok {
		models.SendResponse(w, result, http.StatusBadRequest)
		log.Printf("result: %v\n", result)
		return
	}

	// 既存ユーザーがEmailを使用していないかチェック

	// ユーザーを登録する準備
	var createUser models.User
	createUser.FirstName = user.FirstName
	createUser.LastName = user.LastName
	createUser.Email = user.Email
	createUser.Password = models.EncryptPassword(user.Password)

	fmt.Printf("createUser: %v\n", createUser)
	fmt.Printf("&createUser: %v\n", &createUser)

	// 実際にDBに登録する
	if err := createUser.CreateUser(); err != nil {
		models.SendResponse(w, "Unable to register user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	models.SendResponse(w, "Successfully Signup", http.StatusOK)

}
func signinFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK")
}

func createHabitFunc(w http.ResponseWriter, r *http.Request) {
	// ユーザーを取得する
	DB.Where()
	// 登録する
}
