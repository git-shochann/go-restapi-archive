package models

import (
	"fmt"
	"log"
)

// 値渡し
func (u User) CreateUser() error {

	fmt.Println("CreateUser!")
	// uのポインタ(アドレス)を渡す
	if err := DB.Create(&u).Error; err != nil {
		log.Fatalln(err)
	}
	return nil
}
