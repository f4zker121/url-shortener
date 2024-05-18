package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationsError(errs validator.ValidationErrors) Response {
	var errsMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is not a valid url", err.Field()))
		default:
			errsMsgs = append(errsMsgs, fmt.Sprintf("field %s is not a valid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errsMsgs, ","),
	}
}
