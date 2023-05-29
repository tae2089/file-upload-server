package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tae2089/file-upload-server/api/router"
)

func main() {
	r := gin.New()
	router.Setup(r)
	r.Run(":8080")
}
