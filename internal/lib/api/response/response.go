package response

import ( "fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error omitempty"`
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

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for i := 0; i < len(errs); i++ {
		switch errs[i].ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", errs[i].Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("this url: %s is not valid", errs[i].Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", errs[i].Field()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
