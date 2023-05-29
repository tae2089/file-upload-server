package aws

import (
	"io"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsService interface {
	UploadFileToS3(file io.Reader, fileName string) (string, error)
}

var (
	awsService  AwsService
	awsOnceSync sync.Once
)

func NewAwsService(s3Client *s3.Client, bucket string) AwsService {
	if awsService == nil {
		awsOnceSync.Do(func() {
			awsService = &awsServiceImpl{
				s3Client: s3Client,
				bucket:   bucket,
			}
		})
	}
	return awsService
}
