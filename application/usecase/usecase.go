package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/farkroft/auth-service/application/repository"
	"gitlab.com/farkroft/auth-service/application/request"
	"gitlab.com/farkroft/auth-service/external/config"
	"gitlab.com/farkroft/auth-service/external/redis"
)

// CaseRepo interface
type CaseRepo interface {
	UserRegister(req request.UserRequest) (int, string, error)
	UserLogin(req request.UserRequest) (int, string, interface{}, error)
	UserAuthVerify(str string) (*jwt.Token, error)
}

// UseCase given ...
type UseCase struct {
	UserRepo repository.UserRepository
	Cfg      config.Repository
	Redis    redis.Repository
}

// NewUsecase return data source instance
func NewUsecase(repo repository.UserRepository, cfg config.Repository, rd redis.Repository) *UseCase {
	return &UseCase{
		UserRepo: repo,
		Cfg:      cfg,
		Redis:    rd,
	}
}
