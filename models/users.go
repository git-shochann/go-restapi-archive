package models

import (
	"fmt"
	"log"
)

func (u *User) CreateUser() error {

	fmt.Println("CreateUser!")
	// 戻り値のErrorフィールドを参照
	// invalid memory address or nil pointer dereference
	// fmt.Printf("u: %+v\n", u)
	// fmt.Printf("*u: %+v\n", *u) // アドレスの値を参照する -> デリファレンス
	if err := DB.Create(&u).Error; err != nil {
		log.Fatalln(err)
	}
	return nil
}
