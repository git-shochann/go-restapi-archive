package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 100%ではなくまずは完了させることを目指す！
func CreateHabitFunc(w http.ResponseWriter, r *http.Request) {

	// Bodyを検証
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	var habitVaridation models.CreateHabitVaridation
	err = json.Unmarshal(reqBody, &habitVaridation)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	fmt.Println(habitVaridation) // Hello

	// JWTの検証
	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	// JWTにIDが乗っているので、IDをもとに保存処理をする

	var habit models.Habit
	habit.Content = habitVaridation.Content
	habit.UserID = userID

	err = habit.CreateHabit()
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}
	fmt.Printf("habit: %v\n", habit) // 時間が含まれていない

	// ここの時点でhabitの実態は書き変わっているはず...。

	// WIP
	response, err := json.Marshal(habit)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	models.SendResponse(w, response, http.StatusOK)
}

func DeteteHabitFunc(w http.ResponseWriter, r *http.Request) {

	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}
	fmt.Println("JWTの検証完了!!")

	// 確認したJWTのクレームのuser_id
	// パスパラメーターから取得する habitのid
	vars := mux.Vars(r)
	fmt.Printf("vars: %v\n", vars) // vars: map[id:1]
	habitIDStr := vars["id"]

	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		// 今だと... strconv.Atoi: parsing "": invalid syntax とただのエラーメッセージが返却される
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	var habit models.Habit

	err = models.DeleteHabit(habitID, userID, habit)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}
	models.SendResponse(w, nil, http.StatusOK)

}

func UpdateHabitFunc(w http.ResponseWriter, r *http.Request) {

	// Bodyを検証
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	var habitVaridation models.CreateHabitVaridation
	err = json.Unmarshal(reqBody, &habitVaridation)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	fmt.Println(habitVaridation)

	// JWTの検証
	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	var habit models.Habit
	habit.Content = habitVaridation.Content
	habit.UserID = userID

	// アップデートに必要なのは、habitid, content, (後ほど...finished)

	err = habit.UpdateHabit()
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	response, err := json.Marshal(habit)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	models.SendResponse(w, response, http.StatusOK)

}

// ユーザー1人が持っているhabitを全て取得する
func GetAllHabitFunc(w http.ResponseWriter, r *http.Request) {

	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}
	fmt.Println("JWTの検証完了!!")

	// パスパラメーターでuserid取得可能
	vars := mux.Vars(r)
	parameterUserIDStr := vars["id"]

	parameterUserID, err := strconv.Atoi(parameterUserIDStr)
	if err != nil {
		// 今だと... strconv.Atoi: parsing "": invalid syntax とただのエラーメッセージが返却される
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	// パスパラメーターのID + JWTで検証したIDが一致しなければエラー
	if userID != parameterUserID {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	var user models.User
	user.ID = uint(userID)

	var habit []models.Habit
	err = user.GetAllHabitByUserID(&habit) // 旧: 値を渡す, 新: ポインタ(アドレス)を渡すことでしっかりと返却された
	if err != nil {
		// 今だと... strconv.Atoi: parsing "": invalid syntax とただのエラーメッセージが返却される
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	response, err := json.Marshal(habit)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	models.SendResponse(w, response, http.StatusOK)
}
