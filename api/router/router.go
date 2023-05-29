package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/file-upload-server/api/middleware"
	"github.com/tae2089/file-upload-server/internal/config"
)

func Setup(r *gin.Engine) {
	if gin.Mode() == "debug" {
		config.LoadEnv()
	}
	fileRouter := r.Group("")
	registFileRouter(fileRouter)
	fileRouter.Use(middleware.JwtAuthMiddleware())
	healthRouter := r.Group("")
	registHealthRouter(healthRouter)
}
