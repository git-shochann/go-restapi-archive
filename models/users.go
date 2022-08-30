package models

// エラーハンドリング
// https://gorm.io/ja_JP/docs/error_handling.html

// 値渡し
func (u User) CreateUser() error {

	// uのポインタ(アドレス)を渡す
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil

}

// Emailを元に重複していないか検索をする
// User構造体の値はなぜ必要？ -> 結果を格納するため(out) -> First(out interface{}, where ...interface{}) *gorm.DB
func GetUserByEmail(u User, email string) error {
	// 検索する
	if err := DB.Where("email = ?", email).First(&u).Error; err != nil {
		return err
	}
	// 結果の値を出力
	// fmt.Printf("user: %v\n", u)
	return nil
}
