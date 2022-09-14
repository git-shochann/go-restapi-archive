package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserSignupValidation struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=15,containsany=0123456789"`
}

func (u UserSignupValidation) SignupValidator() (string, error) {

	validate := validator.New()
	err := validate.Struct(&u)

	fmt.Printf("err: %v\n", err)

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
		return errorMessage, err
	}
	return "", err
}

// ログインのバリデーション
type UserSigninValidation struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15,containsany=0123456789"`
}

func (u UserSigninValidation) SigninValidator() (string, error) {
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
		return errorMessage, err
	}
	return "", err
}

// 習慣を登録するときのバリデーション
type CreateHabitValidation struct {
	Content string `json:"content" validate:"required"`
}

func (c CreateHabitValidation) CreateHabitValidator() (string, error) {
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
		return errorMessage, err
	}
	return "", err
}
