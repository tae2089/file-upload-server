package aws

import (
	"fmt"
	"os"
	"testing"

	"github.com/tae2089/file-upload-server/internal/config"
)

func Test_awsServiceImpl_UploadFileToS3(t *testing.T) {
	t.Skip()
	config.LoadEnv()
	client := config.NewS3Client()
	awsService := &awsServiceImpl{
		s3Client: client,
		bucket:   os.Getenv("BUCKET"),
	}
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	got, err := awsService.UploadFileToS3(file, file.Name())
	fmt.Println(got)
}
