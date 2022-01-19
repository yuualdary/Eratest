package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {

	fmt.Println(data)

	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {

	var errors []string

	for _, checkError := range err.(validator.ValidationErrors) {

		errors = append(errors, checkError.Error())
	}
	return errors
}
