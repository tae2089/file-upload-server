package controller

import (
	"sync"

	"github.com/tae2089/file-upload-server/internal/aws"

	"github.com/gin-gonic/gin"
)

type FileController interface {
	UploadFileToCloud(c *gin.Context)
}

var (
	fileController     FileController
	fileControllerOnce sync.Once
)

func NewFileController(awsService aws.AwsService) FileController {
	if fileController == nil {
		fileControllerOnce.Do(func() {
			fileController = &fileControllerImpl{
				awsService,
			}
		})
	}
	return fileController
}
