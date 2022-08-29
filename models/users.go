package models

// 値渡し
func (u User) CreateUser() error {

	// uのポインタ(アドレス)を渡す
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil

}
