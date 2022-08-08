package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func init() {
	// 環境変数の読み込み
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Enable load env file: %v", err)
	}
}
