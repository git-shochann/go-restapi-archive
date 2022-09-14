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
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// バリデーションの実施
	var habitValidation models.CreateHabitValidation
	err = json.Unmarshal(reqBody, &habitValidation)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	errorMessage, err := habitValidation.CreateHabitValidator()

	if err != nil {
		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		log.Println(err)
		return
	}

	// JWTの検証
	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, "Authentication error", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// JWTにIDが乗っているので、IDをもとに保存処理をする

	var habit models.Habit
	habit.Content = habitValidation.Content
	habit.UserID = userID

	err = habit.CreateHabit()
	if err != nil {
		models.SendErrorResponse(w, "Failed to create habit", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response, err := json.Marshal(habit)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	models.SendResponse(w, response, http.StatusOK)
}

func DeteteHabitFunc(w http.ResponseWriter, r *http.Request) {

	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, "authentication error", http.StatusBadRequest)

		log.Println(err)
		return
	}

	// 確認したJWTのクレームのuser_id
	// パスパラメーターから取得する habitのid
	vars := mux.Vars(r)
	fmt.Printf("vars: %v\n", vars) // vars: map[id:1]
	habitIDStr := vars["id"]

	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		models.SendErrorResponse(w, "Something wrong", http.StatusBadRequest)

		log.Println(err)
		return
	}

	var habit models.Habit

	err = models.DeleteHabit(habitID, userID, &habit)
	if err != nil {
		models.SendErrorResponse(w, "Failed to delete habit", http.StatusBadRequest)

		log.Println(err)
		return
	}
	models.SendResponse(w, nil, http.StatusOK)

}

// WIP: 現在1つのIDを送ってるのにそのユーザーに紐付く習慣全て変わってる
func UpdateHabitFunc(w http.ResponseWriter, r *http.Request) {

	// JWTの検証
	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, "authentication error", http.StatusBadRequest)

		log.Println(err)
		return
	}

	// 確認したJWTのクレームのuser_id
	// パスパラメーターから取得する habitのid

	vars := mux.Vars(r)
	fmt.Printf("vars: %v\n", vars) // vars: map[id:1]
	habitIDStr := vars["id"]

	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		models.SendErrorResponse(w, "Something wrong", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Bodyを検証
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// バリデーションの実施
	var habitValidation models.CreateHabitValidation
	err = json.Unmarshal(reqBody, &habitValidation)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	errorMessage, err := habitValidation.CreateHabitValidator()

	if err != nil {
		models.SendErrorResponse(w, errorMessage, http.StatusBadRequest)
		log.Println(err)
		return
	}

	var habit models.Habit
	habit.Model.ID = uint(habitID)          // id(habit)
	habit.Content = habitValidation.Content // content
	habit.UserID = userID                   // user_id

	err = habit.UpdateHabit()
	if err != nil {
		models.SendErrorResponse(w, "Failed to update habit", http.StatusBadRequest)
		log.Println(err)
		return
	}
	fmt.Printf("habit: %v\n", habit)

	response, err := json.Marshal(habit)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)

		log.Println(err)
		return
	}

	models.SendResponse(w, response, http.StatusOK)

}

// ユーザー1人が持っているhabitを全て取得する
func GetAllHabitFunc(w http.ResponseWriter, r *http.Request) {

	// JWTの検証とユーザーIDの取得
	userID, err := models.CheckJWTToken(r)
	if err != nil {
		models.SendErrorResponse(w, "authentication error", http.StatusBadRequest)

		log.Println(err)
		return
	}

	var user models.User
	user.ID = uint(userID)

	var habit []models.Habit
	err = user.GetAllHabitByUserID(&habit) // 旧: 値を渡す, 新: ポインタ(アドレス)を渡すことでしっかりと返却された
	if err != nil {
		models.SendErrorResponse(w, "Failed to get habit", http.StatusBadRequest)

		log.Println(err)
		return
	}

	response, err := json.Marshal(habit)
	if err != nil {
		models.SendErrorResponse(w, "Failed to read json", http.StatusBadRequest)

		log.Println(err)
		return
	}

	models.SendResponse(w, response, http.StatusOK)
}
