package models

import (
	"fmt"

	"github.com/go-playground/validator"
)

type UserSignupVaridation struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,lowercase,numeric,min=8,max=15"`
}

// こういうところはポインタレシーバーにするべきか？ -> uを参照して変更しないので値レシーバーでOK
// https://qiita.com/Yuuki557/items/e9f5bdfbbfe92973a05e
func (u UserSignupVaridation) SignupVaridator() (ok bool, result string) {

	validate := validator.New()
	err := validate.Struct(u)

	var errorMessage string // nilで埋められる

	if err != nil {
		fmt.Println(err.(validator.ValidationErrors))
		for _, fieldErr := range err.(validator.ValidationErrors) {

			fieldName := fieldErr.Field()

			switch fieldName {
			case "FirstName":
				errorMessage = "Invalid First Name"
			case "LastName":
				errorMessage = "Invalid Last Name"
			case "Email":
				errorMessage = "Invalid Email"
			case "Password":
				errorMessage = "Invalid Password"
			}
		}
		return false, errorMessage
	}
	return true, errorMessage
}
