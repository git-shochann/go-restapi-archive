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

func DeleteHabit(userID, habitID int, habit Habit) error {

	// &habitが必要なのはなぜ？
	if err := DB.Where("id = ?", habitID).Delete(&habit).Error; err != nil {
		return err
	}
	return nil
}
