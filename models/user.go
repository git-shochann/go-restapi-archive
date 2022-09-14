package models

// エラーハンドリング
// https://gorm.io/ja_JP/docs/error_handling.html

// ポインタ渡し -> 元の実体を書き換えるので
func (u *User) CreateUser() error {

	if err := DB.Create(u).Error; err != nil {
		return err
	}
	return nil

}

// Emailを元に重複していないか検索をする
// User構造体の値はなぜ必要？ -> 結果を格納するため(out) -> First(out interface{}, where ...interface{}) *gorm.DB
func GetUserByEmail(u *User, email string) error {
	// fmt.Printf("user1: %v\n", u) // user1: {{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>}    }
	// 検索する
	if err := DB.Where("email = ?", email).First(u).Error; err != nil {
		return err
	}
	// 結果の値を出力
	// fmt.Printf("user2: %v\n", u) // user2: {{1 2022-08-31 08:42:39 +0000 UTC 2022-08-31 08:42:39 +0000 UTC <nil>} test test test@gmail.com $2a$10$9QZC62Z6JODHtR1Kg1WCPuH9dAvbU64XJkJNlSwCIwlpbBQ84Eqxq}
	return nil
}
