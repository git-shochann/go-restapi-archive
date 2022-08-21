package models

import "log"

func (u User) CreateUser() error {
	// 戻り値のErrorフィールドを参照
	if err := DB.Create(&u).Error; err != nil {
		log.Fatalln(err)
	}
	return nil
}
