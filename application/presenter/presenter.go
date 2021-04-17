package presenter

import (
	"fmt"

	"gitlab.com/farkroft/auth-service/application/response"
)

// ErrorPresenter presenter of error response
func ErrorPresenter(str string, err error) response.ErrorResponse {
	strErr := fmt.Sprintf(str+": %s", err.Error())
	return response.ErrorResponse{Message: strErr}
}

// SuccessPresenter presenter of success response
func SuccessPresenter(isSuccess bool, msg string, data interface{}) response.SuccessResponse {
	return response.SuccessResponse{
		Success: isSuccess,
		Message: msg,
		Data:    data,
	}
}
