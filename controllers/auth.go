package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignupFunc(w http.ResponseWriter, r *http.Request) {

	// Jsonでくるので、まずGoで使用できるようにする
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	var signupUser models.UserSignupVaridation
	err = json.Unmarshal(reqBody, &signupUser)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	ok, result := signupUser.SignupVaridator()

	// false時の処理
	if !ok {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Printf("result: %v\n", result)
		return
	}

	// ユーザーを登録する準備
	var createUser models.User
	createUser.FirstName = signupUser.FirstName
	createUser.LastName = signupUser.LastName
	createUser.Email = signupUser.Email
	createUser.Password = models.EncryptPassword(signupUser.Password)

	fmt.Printf("createUser: %v\n", createUser)
	fmt.Printf("&createUser: %v\n", &createUser)

	// 実際にDBに登録する
	if err := createUser.CreateUser(); err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// 重複時
	// err.Error() -> "Error 1062: Duplicate entry 'test@gmail.com' for key 'users.email'"

	// 成功！
	// なぜ登録出来なかったのかもう少し詳細のメッセージがあっていいかも。
	// errを使って "message": err など？
	if err := models.SendAuthResponse(w, &createUser, 200); err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

}

// サインイン(ログイン)
func SigninFunc(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	var signinUser models.UserSigninVaridation
	if err := json.Unmarshal(reqBody, &signinUser); err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// バリデーションの実施
	ok, result := signinUser.SigninVaridator()

	if !ok {
		models.SendErrorResponse(w, result, http.StatusBadRequest)
		log.Println(err)
		return
	}

	// 結果を格納する構造体の生成
	var user models.User

	// emailでユーザーを検索する -> 成功したらuserに値が入る
	// user自体の実態を書き換えるのでアドレスを渡してあげる
	err = models.GetUserByEmail(&user, signinUser.Email)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	// 実態を書き換えたそのままuserを出力する
	fmt.Printf("user: %v\n", user) // {{1 2022-08-31 08:42:39 +0000 UTC 2022-08-31 08:42:39 +0000 UTC <nil>} test test test@gmail.com $2a$10$9QZC62Z6JODHtR1Kg1WCPuH9dAvbU64XJkJNlSwCIwlpbBQ84Eqxq}

	// ここでログインユーザーを取得出来たのでuserを使ってく
	// bcryptでDBはハッシュかしているので比較する関数

	fmt.Printf("signinUser.Password: %v\n", signinUser.Password)
	fmt.Printf("user.Password: %v\n", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signinUser.Password))
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := models.SendAuthResponse(w, &user, 200); err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

}
