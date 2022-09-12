package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func LoggingSetting() {
	file, err := os.OpenFile("logging.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, file)      // 出力先を2つ設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 出力のフォーマットを設定
	log.SetOutput(multiLogFile)                          // 実際に設定を反映

}

func EncryptPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

// ステータスコード200の場合のレスポンス
func SendResponse(w http.ResponseWriter, response []byte, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

// ステータスコード200以外のレスポンスで使用
// message: err.Error() とする
func SendErrorResponse(w http.ResponseWriter, myMessage string, errorMessage string, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := map[string]string{
		"message": myMessage,
		"detail":  errorMessage,
	}
	// jsonに変換する
	responseBody, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = w.Write(responseBody)
	if err != nil {
		return err
	}
	return nil
}

// 新規登録とログイン時のレスポンスとしてJWTトークンとUser構造体を返却する
func SendAuthResponse(w http.ResponseWriter, user *User, code int) error {
	fmt.Println("SendAuthResponse!")

	jwtToken, err := user.CreateJWTToken()
	if err != nil {
		return err
	}

	// レスポンス
	response := AuthResponse{
		User:     *user,
		JwtToken: jwtToken,
	}

	// 構造体をjsonに変換
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	fmt.Printf("jsonResponse: %v\n", string(jsonResponse))

	if err := SendResponse(w, jsonResponse, code); err != nil {
		return err
	}

	return nil

}
