package helpers

import (
	model "GoArticle/app/model"
	database "GoArticle/config"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func ValidateStruct(data interface{}) map[string]string {

	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Tag()
	}
	return errors
}

func uniqueEmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	result := database.DB.Where("email = ?", email).Find(&model.User{})

	if result.Error != nil || result.RowsAffected > 0 {
		return false
	}

	return true
}

func uniqueUsernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	result := database.DB.Where("username = ?", username).Find(&model.User{})
	if result.Error != nil || result.RowsAffected > 0 {
		return false
	}

	return true
}

func RegisterCustomValidations() {
	validate.RegisterValidation("uniqueEmail", uniqueEmailValidation)
	validate.RegisterValidation("uniqueUsername", uniqueUsernameValidation)
}
