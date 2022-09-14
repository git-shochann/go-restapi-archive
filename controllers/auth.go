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
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		return
	}

	var signupUser models.UserSignupValidation
	err = json.Unmarshal(reqBody, &signupUser)

	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	errorMessage, err := signupUser.SignupValidator()
	if err != nil {
		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		log.Println(err)
		return
	}

	// ユーザーを登録する準備
	var createUser models.User
	createUser.FirstName = signupUser.FirstName
	createUser.LastName = signupUser.LastName
	createUser.Email = signupUser.Email
	createUser.Password = models.EncryptPassword(signupUser.Password)

	// 実際にDBに登録する
	if err := createUser.CreateUser(); err != nil {
		models.SendErrorResponse(w, "Faild to create user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// createUser -> ポインタ型(アドレス)
	if err := models.SendAuthResponse(w, &createUser, http.StatusOK); err != nil {
		models.SendErrorResponse(w, "Unknown error occurred", http.StatusBadRequest)
		log.Println(err)
		return
	}

}

// サインイン(ログイン)
func SigninFunc(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	var signinUser models.UserSigninValidation
	if err := json.Unmarshal(reqBody, &signinUser); err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// バリデーションの実施
	errorMessage, err := signinUser.SigninValidator()

	if err != nil {
		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		log.Println(err)
		return
	}

	// 結果を格納する構造体の生成
	var user models.User

	// emailでユーザーを検索する -> 成功したらuserに値が入る
	// user自体の実体を書き換えるのでアドレスを渡してあげる
	err = models.GetUserByEmail(&user, signinUser.Email)
	if err != nil {
		models.SendErrorResponse(w, "Faild to get user", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	// 実体を書き換えたそのままuserを出力する
	fmt.Printf("user: %v\n", user) // {{1 2022-08-31 08:42:39 +0000 UTC 2022-08-31 08:42:39 +0000 UTC <nil>} test test test@gmail.com $2a$10$9QZC62Z6JODHtR1Kg1WCPuH9dAvbU64XJkJNlSwCIwlpbBQ84Eqxq}

	// ここでログインユーザーを取得出来たのでuserを使ってく
	// bcryptでDBはハッシュかしているので比較する関数

	fmt.Printf("signinUser.Password: %v\n", signinUser.Password)
	fmt.Printf("user.Password: %v\n", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signinUser.Password))
	if err != nil {
		models.SendErrorResponse(w, "Password error occurred", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := models.SendAuthResponse(w, &user, 200); err != nil {
		models.SendErrorResponse(w, "Faild to sign in", http.StatusBadRequest)
		log.Println(err)
		return
	}

}
