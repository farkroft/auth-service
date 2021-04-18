package server

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"gitlab.com/farkroft/auth-service/application/controller"
	"gitlab.com/farkroft/auth-service/external/config"
	"gitlab.com/farkroft/auth-service/external/constants"
	// "gitlab.com/farkroft/auth-service/external/log"
)

// NewServer instance of server
func NewServer(cfg *config.Config, ctl *controller.Controller) {
	r := gin.New()
	r.Use(logger.SetLogger())
	NewRouter(r, ctl)
	r.Run(cfg.GetString(constants.EnvPort))
}
