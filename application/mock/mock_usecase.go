package mock

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/farkroft/auth-service/application/request"
)

type MockUseCase struct{}

func (u *MockUseCase) UserRegister(req request.UserRequest) (int, string, error) {
	return http.StatusCreated, "OK", nil
}

func (u *MockUseCase) UserAuthVerify(str string) (int, string, interface{}, error) {
	token := &jwt.Token{}
	return http.StatusOK, "OK", token, nil
}

func (u *MockUseCase) UserLogin(req request.UserRequest) (int, string, interface{}, error) {
	return 0, "", nil, nil
}

func (u *MockUseCase) UserLogout(str string) (int, string, interface{}, error) {
	return http.StatusOK, "OK", nil, nil
}
