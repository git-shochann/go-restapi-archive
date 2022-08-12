package model

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// グローバルに宣言
var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func ConnectDB() {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// test:test@tcp(mysql_db:3306)/test?charset=utf8mb4&parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(mysql_db:3306)/%s?charset=utf8mb4&parseTime=true", user, pass, dbName)

	// コネクションプールの生成
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Enable Connect to DB: %v", err)
	} else {
		fmt.Println("Successfully Connected DB")
	}

	// テーブルの作成
	DB.AutoMigrate(&User{})

}
