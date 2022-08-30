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
	// router.HandleFunc("api/v1/create", createHabitFunc).Methods("POST")
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
	var signupUser models.UserSignupVaridation
	err = json.Unmarshal(reqBody, &signupUser)
	if err != nil {
		models.SendErrorResponse(w, "Unable to unmarshal json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// fmt.Printf("signupUser: %v\n", signupUser)

	ok, result := signupUser.SignupVaridator()

	// false時の処理
	if !ok {
		models.SendErrorResponse(w, result, http.StatusBadRequest)
		log.Printf("result: %v\n", result)
		return
	}

	// 既存ユーザーがEmailを使用していないかチェック
	var user models.User

	// メールアドレスがある -> 登録NG / メールアドレスがない -> 登録OK
	err = models.GetUserByEmail(user, signupUser.Email)
	// BUG: errに値が入り、Record not found が返ってきてしまう -> ただそれはOKとしたい -> ただなんらかのエラーもエラーハンドリングすべき
	if err != nil {
		models.SendErrorResponse(w, "Something wrong", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// userの値があるかどうかでチェックする
	// fmt.Printf("user: %+v\n", user)
	// os.Exit(1)

	// // ErrRecordNotFoundが出ない -> 登録出来ない
	// if !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	models.SendErrorResponse(w, "Email address is already in use", http.StatusBadRequest)
	// 	return
	// }

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
		models.SendErrorResponse(w, "Unable to register user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// 成功！
	if err := models.SendAuthResponse(w, &createUser, 200); err != nil {
		models.SendErrorResponse(w, "Something wrong", http.StatusBadRequest)
		log.Println(err)
		return
	}

}

// サインイン(ログイン)
func signinFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK")
}
