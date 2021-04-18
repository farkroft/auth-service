package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/farkroft/auth-service/application/presenter"
	"gitlab.com/farkroft/auth-service/application/request"
	"gitlab.com/farkroft/auth-service/external/log"
	"gitlab.com/farkroft/auth-service/external/util"
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

	httpCode, strResp, resp, err := ctl.UserUseCase.UserLogin(req)
	if err != nil {
		if util.IsErrorRecordNotFound(err) {
			c.JSON(http.StatusOK, presenter.ErrorPresenter(strResp, err))
			return
		}
		log.Errorf(strResp, err)
		c.JSON(httpCode, presenter.ErrorPresenter(strResp, err))
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessPresenter(true, strResp, resp))
}

// UserAuth vaidate token from header
func (ctl *Controller) UserAuth(c *gin.Context) {
	strToken := c.GetHeader("token")

	httpCode, strResp, resp, err := ctl.UserUseCase.UserAuthVerify(strToken)
	if err != nil {
		log.Errorf("token verify", err)
		c.JSON(httpCode, presenter.ErrorPresenter(strResp, err))
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessPresenter(true, strResp, resp))
}

// Logout account
func (ctl *Controller) Logout(c *gin.Context) {
	strToken := c.GetHeader("token")

	httpCode, strResp, resp, err := ctl.UserUseCase.UserLogout(strToken)
	if err != nil {
		if util.IsErrorRecordNotFound(err) {
			c.JSON(http.StatusOK, presenter.ErrorPresenter(strResp, err))
			return
		}
		log.Errorf(strResp, err)
		c.JSON(httpCode, presenter.ErrorPresenter(strResp, err))
		return
	}

	c.JSON(http.StatusOK, presenter.SuccessPresenter(true, strResp, resp))
}
