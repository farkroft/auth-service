package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/farkroft/auth-service/application/controller"
)

// NewRouter router
func NewRouter(r *gin.Engine, ctl *controller.Controller) {
	r.GET("/ping", ctl.Ping)
	r.POST("/register", ctl.Register)
	r.POST("/login", ctl.Login)
	r.POST("/verify-auth", ctl.UserAuth)
}
