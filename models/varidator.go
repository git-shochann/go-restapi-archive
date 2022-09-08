package models

import (
	"github.com/go-playground/validator/v10"
)

type UserSignupVaridation struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=15,numeric,lowercase"` // TODO: Error:Field validation for 'Password' failed on the 'numeric' tag
}

func (u UserSignupVaridation) SignupVaridator() (ok bool, errMessage string) {

	validate := validator.New()
	err := validate.Struct(&u)

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

// ログインのバリデーション
type UserSigninVaridation struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15,lowercase"` // TODO: 上記と同様
}

func (u UserSigninVaridation) SigninVaridator() (ok bool, errMessage string) {
	validate := validator.New()
	err := validate.Struct(&u)

	var errorMessage string

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			fieldName := fieldErr.Field()

			switch fieldName {
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

// 習慣を登録するときのバリデーション
type CreateHabitVaridation struct {
	Content string `json:"content" validate:"required"`
}

func (c CreateHabitVaridation) CreateHabitVaridator() (ok bool, errMessage string) {
	validate := validator.New()
	err := validate.Struct(&c)

	var errorMessage string

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			fieldName := fieldErr.Field()

			switch fieldName {
			case "Content":
				errorMessage = "Invalid Content"

			}
		}
		return false, errorMessage
	}
	return true, errorMessage
}
