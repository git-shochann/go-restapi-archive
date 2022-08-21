package models

import (
	"fmt"

	"github.com/go-playground/validator"
)

// TODO:
type UserSignupVaridation struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" varidate:"required"`
	Email     string `json:"email" varidate:"required,email"`
	Password  string `json:"password" varidate:"required,lowercase,numeric,min=8,max=15"`
}

// こういうところはポインタレシーバーにするべきか？ -> uを参照して変更しないので値レシーバーでOK
// https://qiita.com/Yuuki557/items/e9f5bdfbbfe92973a05e
func (u UserSignupVaridation) SignupVaridator() (ok bool, result map[string]string) {

	validate := validator.New()
	err := validate.Struct(u)

	errorMessage := make(map[string]string) // nilで埋められる

	if err != nil {
		fmt.Println(err.(validator.ValidationErrors))
		for _, fieldErr := range err.(validator.ValidationErrors) {

			fieldName := fieldErr.Field()

			switch fieldName {
			case "FirstName":
				errorMessage["FirstName"] = "Invalid First Name"
			case "LastName":
				errorMessage["LastName"] = "Invalid Last Name"
			case "Email":
				errorMessage["Email"] = "Invalid Email"
			case "Password":
				errorMessage["Password"] = "Invalid Password"
			}
		}
		return false, errorMessage
	}
	return true, errorMessage
}
