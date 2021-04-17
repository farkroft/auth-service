package controller

import (
	"gitlab.com/farkroft/auth-service/application/usecase"
)

// Controller struct
type Controller struct {
	UserUseCase usecase.CaseRepo
}

// NewController func
func NewController(u *usecase.UseCase) *Controller {
	return &Controller{UserUseCase: u}
}
