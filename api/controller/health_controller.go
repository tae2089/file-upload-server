package controller

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	CheckHealth(c *gin.Context)
}

var (
	healthController     HealthController
	healthControllerOnce sync.Once
)

func NewHealthController() HealthController {
	if healthController == nil {
		healthControllerOnce.Do(func() {
			healthController = &healthControllerImpl{}
		})
	}
	return healthController
}
