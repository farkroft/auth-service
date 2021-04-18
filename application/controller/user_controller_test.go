package controller_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/farkroft/auth-service/application/controller"
	"gitlab.com/farkroft/auth-service/application/mock"
	"gitlab.com/farkroft/auth-service/application/request"
	"gitlab.com/farkroft/auth-service/application/response"
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

func TestUserLoginShouldSuccessAndReturn200(t *testing.T) {
	mock.ServerMock(func(r *gin.Engine) {
		m := new(mock.MockUseCase)
		ctl := controller.Controller{UserUseCase: m}
		r.POST("/login", ctl.Login)

		req := request.UserRequest{
			Username: "fajarar77@gmail.com",
			Password: "password",
		}

		data, _ := json.Marshal(req)
		body := bytes.NewReader(data)
		request, _ := http.NewRequest(http.MethodPost, "/login", body)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, request)
		bytes, _ := ioutil.ReadAll(res.Body)
		successResponse := response.SuccessResponse{}
		expectedResult := response.SuccessResponse{
			Success: true,
			Message: "",
			Data:    nil,
		}
		_ = json.Unmarshal(bytes, &successResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Exactly(t, successResponse, expectedResult)
	})
}

func TestUserAuthShouldSuccessAndReturn200(t *testing.T) {
	mock.ServerMock(func(r *gin.Engine) {
		m := new(mock.MockUseCase)
		ctl := controller.Controller{UserUseCase: m}
		r.POST("/verify-auth", ctl.UserAuth)

		request, _ := http.NewRequest(http.MethodPost, "/verify-auth", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, request)
		bytes, _ := ioutil.ReadAll(res.Body)
		successResponse := response.SuccessResponse{}
		expectedResult := response.SuccessResponse{
			Success: true,
			Message: "OK",
			Data:    nil,
		}
		_ = json.Unmarshal(bytes, &successResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Exactly(t, successResponse, expectedResult)
	})
}

func TestLogoutShouldSuccessAndReturn200(t *testing.T) {
	mock.ServerMock(func(r *gin.Engine) {
		m := new(mock.MockUseCase)
		ctl := controller.Controller{UserUseCase: m}
		r.POST("/logout", ctl.Logout)

		request, _ := http.NewRequest(http.MethodPost, "/logout", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, request)
		bytes, _ := ioutil.ReadAll(res.Body)
		successResponse := response.SuccessResponse{}
		expectedResult := response.SuccessResponse{
			Success: true,
			Message: "OK",
			Data:    nil,
		}
		_ = json.Unmarshal(bytes, &successResponse)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Exactly(t, successResponse, expectedResult)
	})
}
