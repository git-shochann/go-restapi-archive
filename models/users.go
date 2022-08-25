package models

import (
	"fmt"
	"log"
	"net/http"
)

// 値渡し
func (u User) CreateUser(w http.ResponseWriter) error {

	fmt.Println("CreateUser!")
	fmt.Printf("u: %+v\n", u)   // 値 ->
	fmt.Printf("&u: %+v\n", &u) // ポインタ ->

	// uのポインタ(アドレス)を渡す
	if err := DB.Create(&u).Error; err != nil {
		SendResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return err
	}
	return nil
}
