package models

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthResponse struct {
	User     User // 埋め込む
	JwtToken string
}

// 新規登録が成功したらトークンを発行してレスポンスに含める。
// Userと紐づいているのでメソッドでOK。
func (u *User) CreateJWTToken() (string, error) {

	// クレームの作成
	claim := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダー部分とペイロードの作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 署名をして完全なjwtを生成する
	// 引数にtoken.SignedString(os.Getenv("JWTSIGNKEY")) だとエラー
	jwtToken, err := token.SignedString([]byte(os.Getenv("JWTSIGNKEY")))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

// これでも書ける
func (u *User) WIP() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// func New(method SigningMethod) *Token {
	//	return NewWithClaims(method, MapClaims{}) // ここでMapClaims{}が返ってくる
	// }

	// type Token struct {

	// インターフェース型 interface{} Claimsインターフェース これも型なので実際の型は持っていない
	// 	Claims Claims
	// }

	// type Claims interface {
	//	Valid() error
	// }

	// type MapClaims map[string]interface{}

	// 型アサーションをする　interface型 -> 元の型にしてあげる
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	jwtToken, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	if err != nil {
		return "", err
	}
	fmt.Printf("jwt: %v\n", jwtToken)

	return jwtToken, nil

}

// リクエスト時のJWTTokenの検証
func CheckJWTToken(r *http.Request) {

	// リクエスト構造体を渡す -> リクエストヘッダーの取得する -> Header map[string][]string

	bearerTokenStr := r.Header["Authorization"][0]
	fmt.Printf("token: %#v\n", bearerTokenStr)                    // token: "Bearer jifdaslkjhdafskjhksdfhakfdk"
	token := strings.Split(bearerTokenStr, "Bearer ")             // 第二引数: 何で分割したいのか？
	fmt.Println(token)                                            // jifdaslkjhdafskjhksdfhakfdk
	jwt.Parse(token[0], func(t *jwt.Token) (interface{}, error) { // 第二引数に関数 -> 引数に *jwt.Token型で戻り値を取

	})

}
