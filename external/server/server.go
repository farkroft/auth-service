package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/farkroft/auth-service/application/controller"
	"gitlab.com/farkroft/auth-service/external/config"
	"gitlab.com/farkroft/auth-service/external/constants"
)

// NewServer instance of server
func NewServer(cfg *config.Config, ctl *controller.Controller) {
	r := gin.New()
	NewRouter(r, ctl)
	r.Run(cfg.GetString(constants.EnvPort))
}
