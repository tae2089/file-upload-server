package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/file-upload-server/di"
)

func registHealthRouter(group *gin.RouterGroup) {
	healthController := di.InitializeHealthController()
	group.GET("/healthz", healthController.CheckHealth)
}
