// package models

// import (
// 	"errors"
// )

// func (h *Habit) CreateHabit() error {

// 	if err := DB.Create(h).Error; err != nil {
// 		return err
// 	}
// 	// fmt.Printf("h: %v\n", h) // h: {{2 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 2022-09-07 13:47:28.774095 +0900 JST m=+3.267163626 <nil>} hello false 1}
// 	return nil

// }

// func DeleteHabit(habitID, userID int, habit *Habit) error {

// 	// &habitでもhabitでも問題がないのは内部でリフレクションが行われているため
// 	result := DB.Where("id = ? AND user_id = ?", habitID, userID).Delete(habit)

// 	if err := result.Error; err != nil {
// 		return err
// 	}

// 	// 実際にレコードが存在し、削除されたかどうかの判定は以下で行う
// 	if result.RowsAffected < 1 {
// 		err := errors.New("not found record")
// 		return err
// 	}

// 	return nil
// }

// func (h *Habit) UpdateHabit() error {

// 	result := DB.Model(h).Where("id = ? AND user_id = ?", h.Model.ID, h.UserID).Update("content", h.Content)

// 	if err := result.Error; err != nil {
// 		return err
// 	}

// 	// 実際にレコードが存在し、更新されたかどうかの判定は以下で行う
// 	if result.RowsAffected < 1 {
// 		err := errors.New("not found record") // 当たり前のように論理削除していたら更新は不可
// 		return err
// 	}

// 	return nil
// }

// //実体を受け取って、実体を書き換えるので、戻り値に指定する必要はない。
// // 旧: 値渡し, 新: ポインタを受け取る！s
// func (u User) GetAllHabitByUserID(habit *[]Habit) error {
// 	// habitテーブル内の外部キーであるuseridで全てを取得する
// 	// fmt.Printf("u.ID: %v\n", u.ID)     // 1
// 	// fmt.Printf("[]habit: %v\n", habit) // 空の構造体

// 	// 全て取得したい
// 	if err := DB.Where("user_id = ?", u.ID).Find(habit).Error; err != nil {
// 		// ここの戻り値
// 		return err
// 	}
// 	// fmt.Printf("habit: %v\n", habit) // habit: [{{2 2022-09-07 04:47:29 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is test false 1} {{3 2022-09-07 04:49:30 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} aaa false 1} {{4 2022-09-07 04:49:31 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is test false 1} {{5 2022-09-07 04:50:22 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} This is testbbb false 1} {{6 2022-09-07 04:55:55 +0000 UTC 2022-09-07 07:23:22 +0000 UTC <nil>} aaadsadsa false 1}]
// 	return nil
// }
