package main

import (
	"go-rest-api/controllers"
)

// func loadENV() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Unable to load env file: %v", err)
// 	}

// 	fmt.Println(os.Getenv("SAMPLE_MESSAGE"))
// }

func main() {

	// 	// 環境変数の読み込み
	// 	loadENV()

	// 	// ログ関連の設定
	// 	models.LoggingSetting()

	// 	// DBに接続してテーブルを作成する
	// 	models.ConnectDB()

	// 	// APIサーバーのスタート
	controllers.StartServer()
}
