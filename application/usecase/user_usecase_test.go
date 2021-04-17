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
