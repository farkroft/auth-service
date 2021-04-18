package server

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gitlab.com/farkroft/auth-service/application/controller"
)

// NewRouter router
func NewRouter(r *gin.Engine, ctl *controller.Controller) {
	r.GET("/ping", ctl.Ping)
	r.POST("/register", ctl.Register)
	r.POST("/login", ctl.Login)
	r.POST("/verify-auth", ctl.UserAuth)
	r.POST("/logout", ctl.Logout)

	pprof.Register(r)
}
