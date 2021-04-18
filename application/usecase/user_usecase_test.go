package usecase_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/farkroft/auth-service/application/mock"
	"gitlab.com/farkroft/auth-service/application/request"
	"gitlab.com/farkroft/auth-service/application/usecase"
)

func TestUserRegisterUseCaseShouldSuccessAndReturn201(t *testing.T) {
	mockRepo := new(mock.MockRepository)
	mockConfig := new(mock.MockConfig)
	mockUsecase := usecase.UseCase{
		UserRepo: mockRepo,
		Cfg:      mockConfig,
	}

	userReq := request.UserRequest{
		Username: "fajarar77@gmail.com",
		Password: "password",
	}

	httpCode, msg, err := mockUsecase.UserRegister(userReq)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, httpCode)
	assert.Equal(t, "OK", msg)
}

func TestUserLoginUseCaseShouldSuccessAndReturn201(t *testing.T) {
	mockRepo := new(mock.MockRepository)
	mockConfig := new(mock.MockConfig)
	mockRedis := new(mock.MockRedis)
	uc := usecase.UseCase{
		UserRepo: mockRepo,
		Cfg:      mockConfig,
		Redis:    mockRedis,
	}

	userReq := request.UserRequest{
		Username: "fajarar77@gmail.com",
		Password: "password",
	}

	httpCode, msg, _, err := uc.UserLogin(userReq)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpCode)
	assert.Equal(t, "OK", msg)
}

func TestUserAuthVerifyUseCaseShouldSuccessAndReturn201(t *testing.T) {
	mockRepo := new(mock.MockRepository)
	mockConfig := new(mock.MockConfigSecretKey)
	mockRedis := new(mock.MockRedisSecret)
	uc := usecase.UseCase{
		UserRepo: mockRepo,
		Cfg:      mockConfig,
		Redis:    mockRedis,
	}

	httpCode, msg, _, err := uc.UserAuthVerify("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MTg4MTk5OTN9LCJVc2VySUQiOiIyMmIwOGU5YS1kZTI4LTRmN2ItODQ1My1hZjA2NjM5M2FmMzYiLCJVc2VybmFtZSI6ImZhamFyYXI3N0BnbWFpbC5jb20ifQ.Wwzljw_el6Dc6wNw8Bn6SnRFbSkT6ZYjbumTULPyuYo")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpCode)
	assert.Equal(t, "OK", msg)
}

func TestUserLogoutUseCaseShouldSuccessAndReturn201(t *testing.T) {
	mockRepo := new(mock.MockRepository)
	mockConfig := new(mock.MockConfigSecretKey)
	mockRedis := new(mock.MockRedisSecret)
	uc := usecase.UseCase{
		UserRepo: mockRepo,
		Cfg:      mockConfig,
		Redis:    mockRedis,
	}

	httpCode, msg, _, err := uc.UserLogout("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MTg4MTk5OTN9LCJVc2VySUQiOiIyMmIwOGU5YS1kZTI4LTRmN2ItODQ1My1hZjA2NjM5M2FmMzYiLCJVc2VybmFtZSI6ImZhamFyYXI3N0BnbWFpbC5jb20ifQ.Wwzljw_el6Dc6wNw8Bn6SnRFbSkT6ZYjbumTULPyuYo")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpCode)
	assert.Equal(t, "OK", msg)
}
