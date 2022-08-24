package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserSignupVaridation struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=15"` // 8-15文字, 小文字, 数値 // test000test ??
}

func (u *UserSignupVaridation) SignupVaridator() (ok bool, errMessage string) {

	validate := validator.New()
	err := validate.Struct(&u)

	fmt.Printf("err.(validator.ValidationErrors): %v\n", err.(validator.ValidationErrors))

	var errorMessage string

	if err != nil {

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
