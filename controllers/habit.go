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
	// fmt.Println("JWTの検証完了!!")

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
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}

	var habit models.Habit

	err = models.DeleteHabit(userID, habitID, habit)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		fmt.Println("エラー！")
		log.Println(err)
		return
	}
	models.SendResponse(w, nil, http.StatusOK)

}
