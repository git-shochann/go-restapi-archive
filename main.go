package main

import (
	"fmt"
	"go-rest-api/controllers"
	"go-rest-api/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Enable load env file: %v", err)
	}

	fmt.Println(os.Getenv("SAMPLE_MESSAGE"))
}

func main() {

	// 環境変数の読み込み
	loadENV()

	// DBに接続してテーブルを作成する
	db := model.ConnectDB()

	// サーバーのスタート
	controllers.StartServer()
}
