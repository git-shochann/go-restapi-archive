package models

import (
	"fmt"
)

func (h Habit) CreateHabit() error {
	// ポインタを渡して実態を書き換える
	if err := DB.Create(&h).Error; err != nil {
		return err
	}
	fmt.Printf("h: %v\n", h) // h: {{2 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 <nil>} hello false 1}
	return nil

}

// WIP
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
