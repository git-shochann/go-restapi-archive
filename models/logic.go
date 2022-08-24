package models

import (
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

// jsonメッセージとステータスコードを返却する
func SendResponse(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(message))
}
