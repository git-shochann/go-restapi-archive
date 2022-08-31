package controllers

import (
	"encoding/json"
	"go-rest-api/models"
	"io/ioutil"
	"log"
	"net/http"
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
	var habit models.CreateHabitVaridation
	err = json.Unmarshal(reqBody, &habit)
	if err != nil {
		models.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// JWTの検証
	models.CheckJWTToken(r)

}
