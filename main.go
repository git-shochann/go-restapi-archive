package main

import (
	"fmt"
	"go-rest-api/controllers"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Enable load env file: %v", err)
	}

	fmt.Println(os.Getenv("SAMPLE_MESSAGE"))
}

func connectDB() {
	db, err := gorm.Open("mysql")
	if err != nil {
		log.Fatalf("Enable Connect to DB: %v", err)
	} else {
		fmt.Println("Successfully connect DB")
	}
}

func main() {

	// 環境変数の読み込み
	LoadENV()

	controllers.StartServer()
}
