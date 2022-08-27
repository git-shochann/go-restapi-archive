package models

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// 新規登録が成功したらトークンを発行してレスポンスに含める。

func CreateJWTToken() {
	// ヘッダー部分の作成
	token := jwt.New(jwt.SigningMethodES256)
	fmt.Printf("token: %v\n", token)

	// ペイロードの作成
	token.Claims := jwt.MapClaims{
	}

	// 署名を行う

}
