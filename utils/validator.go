package utils

import "github.com/go-playground/validator"

var CustomValidator *validator.Validate

func init() {
	CustomValidator = validator.New()
}

func CustomValidation(s interface{}) map[string]string {
	if err := CustomValidator.Struct(s); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		listError := make(map[string]string, len(validationErrors))
		for _, validationError := range validationErrors {
			listError[validationError.Field()] = msgForTag(validationError.Tag())
		}
		return listError
	}
	return nil
}

func msgForTag(tag string) string {
	switch tag {
		case "required":
				return "This field is required"
		case "email":
				return "Invalid email"
	}
	return ""
}