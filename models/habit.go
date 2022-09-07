package models

import (
	"fmt"
)

func (h Habit) CreateHabit() error {
	// ポインタを渡して実体を書き換える
	if err := DB.Create(&h).Error; err != nil {
		return err
	}
	fmt.Printf("h: %v\n", h) // h: {{2 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 <nil>} hello false 1}
	return nil

}

// WIP!
func DeleteHabit(habitID, userID int, habit Habit) error {

	// &habitが必要なのはなぜ？
	// 現在論理削除されているのに、再度削除処理が出来てしまう
	if err := DB.Where("id = ? ", habitID).Delete(&habit).Error; err != nil {
		return err
	}
	return nil
}

func (h Habit) UpdateHabit() error {
	if err := DB.Model(&h).Update("content", h.Content).Error; err != nil {
		return err
	}
	return nil
}

//実体を受け取って、実体を書き換えるので、戻り値に指定する必要はない。
// 旧: 値渡し, 新: ポインタを受け取る！
func (u User) GetAllHabitByUserID(habit *[]Habit) error {
	// habitテーブル内の外部キーであるuseridで全てを取得する
	// fmt.Printf("u.ID: %v\n", u.ID)     // 1
	// fmt.Printf("[]habit: %v\n", habit) // 空の構造体

	// 全て取得したい
	if err := DB.Where("user_id = ?", u.ID).Find(habit).Error; err != nil {
		// ここの戻り値
		return err
	}
	// ちゃんと返ってきてる
	fmt.Printf("habit: %v\n", habit) // habit: [{{2 2022-09-07 04:47:29 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is test false 1} {{3 2022-09-07 04:49:30 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} aaa false 1} {{4 2022-09-07 04:49:31 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is test false 1} {{5 2022-09-07 04:50:22 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is testbbb false 1} {{6 2022-09-07 04:55:55 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} aaadsadsa false 1}]
	return nil
}
