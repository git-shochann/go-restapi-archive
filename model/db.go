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

// User has many Habit
type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt を作成
	Name       string `gorm:"not null"`
	Email      string `gorm:"not null"`
	Password   string `gorm:"not null"`
}

// Habit belongs to User
type Habit struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  string `gorm:"not null"`
}

func ConnectDB() *gorm.DB {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// test:test@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=true", user, pass, dbName)

	// コネクションプールの生成
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Enable Connect to DB: %v", err)
	} else {
		fmt.Println("Successfully Connected DB")
	}

	// テーブルの作成
	return DB.AutoMigrate(&User{}, &Habit{})

}
