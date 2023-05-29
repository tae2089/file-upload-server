package di

import (
	"fmt"
	"log"
	"os"

	"github.com/tae2089/file-upload-server/api/controller"
	"github.com/tae2089/file-upload-server/internal/aws"
	"github.com/tae2089/file-upload-server/internal/config"
)

func initializeAwsService() aws.AwsService {
	s3Client := config.NewS3Client()
	bucket := os.Getenv("BUCKET")
	if bucket == "" {
		log.Panic(fmt.Errorf("bucket is not set"))
	}
	return aws.NewAwsService(s3Client, bucket)
}

func InitializeFileController() controller.FileController {
	awsService := initializeAwsService()
	fileController := controller.NewFileController(awsService)
	return fileController
}

func InitializeHealthController() controller.HealthController {
	return controller.NewHealthController()
}
