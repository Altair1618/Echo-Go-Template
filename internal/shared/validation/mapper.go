package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func formatFieldName(field string) string {
	if len(field) == 0 {
		return field
	}

	var result string
	if field[0] >= 'A' && field[0] <= 'Z' {
		result = string(field[0] + 32)
	} else {
		result = string(field[0])
	}

	for i := 1; i < len(field); i++ {
		if field[i] >= 'A' && field[i] <= 'Z' {
			result += "_" + string(field[i]+32)
		} else {
			result += string(field[i])
		}
	}

	return result
}

func formatValidatorError(err validator.FieldError, field string) string {
	switch err.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "uuid":
		return field + " must be a valid UUID"
	case "url":
		return field + " must be a valid URL"
	case "min":
		return field + " must be at least " + err.Param() + " characters long"
	case "max":
		return field + " must be at most " + err.Param() + " characters long"
	case "len":
		return field + " must be exactly " + err.Param() + " characters long"
	case "oneof":
		return field + " must be one of the following values: " + err.Param()
	default:
		return field + " is invalid"
	}
}

func MapValidatorError(err error) map[string]string {
	if err == nil {
		return nil
	}

	validationErrors := make(map[string]string)

	var valErr validator.ValidationErrors
	if errors.As(err, &valErr) {
		for _, fieldErr := range valErr {
			fieldName := formatFieldName(fieldErr.Field())
			validationErrors[fieldName] = formatValidatorError(fieldErr, fieldName)
		}
	} else {
		validationErrors["error"] = "Invalid input format"
	}

	return validationErrors
}
