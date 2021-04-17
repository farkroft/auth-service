package controller

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/farkroft/auth-service/external/util"

	"gitlab.com/farkroft/auth-service/application/presenter"

	"github.com/gin-gonic/gin"
	"gitlab.com/farkroft/auth-service/application/request"
	"gitlab.com/farkroft/auth-service/external/log"
)

// Register new account
func (ctl *Controller) Register(c *gin.Context) {
	req := request.UserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("bind json", err)
		c.JSON(http.StatusInternalServerError, presenter.ErrorPresenter("bind json", err))
		return
	}

	httpCode, strErr, err := ctl.UserUseCase.UserRegister(req)
	if err != nil {
		log.Errorf(strErr, err)
		c.JSON(httpCode, presenter.ErrorPresenter(strErr, err))
		return
	}

	c.JSON(http.StatusCreated, presenter.SuccessPresenter(true, "OK", "Register succeed"))
}

// Login account
func (ctl *Controller) Login(c *gin.Context) {
	req := request.UserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("bind json", err)
		c.JSON(http.StatusInternalServerError, presenter.ErrorPresenter("bind json", err))
		return
	}

	httpCode, strErr, resp, err := ctl.UserUseCase.UserLogin(req)
	if err != nil {
		if util.IsErrorRecordNotFound(err) {
			c.JSON(http.StatusOK, presenter.ErrorPresenter(strErr, err))
			return
		}
		log.Errorf(strErr, err)
		c.JSON(httpCode, presenter.ErrorPresenter(strErr, err))
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessPresenter(true, "OK", resp))
}

// UserAuth vaidate token from header
func (ctl *Controller) UserAuth(c *gin.Context) {
	strToken := c.GetHeader("token")
	extractedToken, err := extractToken(strToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, presenter.ErrorPresenter("token format", err))
		return
	}

	token, err := ctl.UserUseCase.UserAuthVerify(extractedToken)
	if err != nil {
		log.Errorf("token verify", err)
		c.JSON(http.StatusUnauthorized, presenter.ErrorPresenter("token verify", err))
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessPresenter(true, "token is valid", token))
}

func extractToken(token string) (string, error) {
	arrToken := strings.Split(token, " ")
	if len(arrToken) < 2 {
		err := fmt.Errorf("user unauthorized")
		return "", err
	}

	if !strings.Contains(arrToken[0], "Bearer") {
		err := fmt.Errorf("user unauthorized")
		return "", err
	}

	return arrToken[1], nil
}
