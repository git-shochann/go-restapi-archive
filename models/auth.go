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
	splitToken := strings.Split(bearerTokenStr, "Bearer ") // 第二引数: 何で分割したいのかで処理を行う 分割した結果をスライスの中に突っ込む

	// stringのスライス
	// ["","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjI1NDU1OTYsInVzZXJfaWQiOjF9.B1PMPvxl4aGDaJXwXvZPJXxluh5S4lmiq5oen1KWiaU"]
	fmt.Printf("splitToken: %v\n", splitToken)

	// ここのtoken -> 無名関数である(あくまで関数の定義) -> Parse()の内部処理で使用する -> tokenの値を使用可能 -> 関数の説明をしっかり読めば分かる
	parsedToken, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {

		// 型アサーション -> algの検証を行う
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("signature method invalid")
			return nil, err
		}

		// 暗号鍵を返さなくてないいけないとドキュメントに書いてある。SigningMethodHMACのキーは[]byteで返してあげる
		return []byte(os.Getenv("JWTSIGNKEY")), nil

	})

	fmt.Printf("type and value: parsedToken: %+T, %+v\n", parsedToken.Claims, parsedToken.Claims) // -> Claimsインターフェース -> 元の型 -> // map[exp:1.662545596e+09 user_id:1] -> {"user_id":1}

	// 何らかのエラー
	if err != nil {
		return nil, err // WIP
	}

	// これは？
	if !parsedToken.Valid {
		err := errors.New("invalid token")
		return nil, err
	}

	// user_idを取り出したい

	return parsedToken, nil
}
