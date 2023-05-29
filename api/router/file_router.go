package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/file-upload-server/di"
)

func registFileRouter(group *gin.RouterGroup) {
	fileController := di.InitializeFileController()
	group.POST("/file", fileController.UploadFileToCloud)
}
