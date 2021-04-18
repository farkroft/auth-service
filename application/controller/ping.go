package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.com/farkroft/auth-service/external/log"
)

// Ping controller
func (ctl *Controller) Ping(c *gin.Context) {
	log.Infoln("string")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
