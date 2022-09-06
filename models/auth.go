package models

import (
	"errors"
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
func CheckJWTToken(r *http.Request) (*jwt.Token, error) {

	// リクエスト構造体を渡す -> リクエストヘッダーの取得する -> Header map[string][]string

	// Authorizationヘッダーにあるかどうか確認
	bearerTokenStr := r.Header.Get("Authorization")
	if bearerTokenStr == "" {
		err := errors.New("missing token") // errorインターフェースの作成
		return nil, err
	}

	fmt.Printf("token: %#v\n", bearerTokenStr)             // token: "Bearer jifdaslkjhdafskjhksdfhakfdk"
	tokenSlice := strings.Split(bearerTokenStr, "Bearer ") // 第二引数: 何で分割したいのかで処理を行う

	// ここのtokenはどこで取れる？
	parsedToken, err := jwt.Parse(tokenSlice[0], func(token *jwt.Token) (interface{}, error) {

		// 型アサーション -> algの検証を行う
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			err := errors.New("signature method invalid")
			return nil, err
		}

		return os.Getenv("JWTSIGNKEY"), nil

	})

	fmt.Printf("parsedToken: %v\n", parsedToken)

	// 何らかのエラー
	if err != nil {
		err := errors.New("something wrong")
		return nil, err
	}

	// これは？
	if !parsedToken.Valid {
		err := errors.New("invalid token")
		return nil, err
	}

	return parsedToken, nil
}
