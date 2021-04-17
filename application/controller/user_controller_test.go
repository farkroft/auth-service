package controller_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/farkroft/auth-service/application/response"

	"gitlab.com/farkroft/auth-service/application/request"

	"gitlab.com/farkroft/auth-service/application/controller"

	"github.com/gin-gonic/gin"
	"gitlab.com/farkroft/auth-service/application/mock"
)

func TestRegisterUserShouldSuccessAndReturn200(t *testing.T) {
	mock.ServerMock(func(r *gin.Engine) {
		m := new(mock.MockUseCase)
		ctl := controller.Controller{UserUseCase: m}
		r.POST("/register", ctl.Register)

		req := request.UserRequest{
			Username: "fajarar77@gmail.com",
			Password: "password",
		}

		data, _ := json.Marshal(req)
		body := bytes.NewReader(data)
		request, _ := http.NewRequest(http.MethodPost, "/register", body)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, request)
		bytes, _ := ioutil.ReadAll(res.Body)
		successResponse := response.SuccessResponse{}
		expectedResult := response.SuccessResponse{
			Success: true,
			Message: "OK",
			Data:    "Register succeed",
		}
		_ = json.Unmarshal(bytes, &successResponse)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Exactly(t, successResponse, expectedResult)
	})
}
