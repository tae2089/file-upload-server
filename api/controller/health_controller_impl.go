package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthControllerImpl struct {
}

var _ HealthController = (*healthControllerImpl)(nil)

// CheckHealth implements HealthController
func (*healthControllerImpl) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
