package utils

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func APIResponse(code int, msg string, data interface{}) *Response {
	var response = &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}

	return response
}

func SplitErrorInformation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
