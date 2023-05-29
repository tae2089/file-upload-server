package controller

import (
	"net/http"

	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/file-upload-server/internal/aws"

	"github.com/gin-gonic/gin"
)

type fileControllerImpl struct {
	aws.AwsService
}

var _ FileController = (*fileControllerImpl)(nil)

func (controller *fileControllerImpl) UploadFileToCloud(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileReader, err := file.Open()
	defer fileReader.Close()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url, err := controller.UploadFileToS3(fileReader, file.Filename)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"url": url})
}
