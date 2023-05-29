# FILE UPLOAD SERVER

## Description

This repository is a file upload server using golang gin-gonic backend framwork.

## Usage

Usage in the local environment:

Copy the .env.development file and create a new file called .env.
Execute the ras.sh script in the script folder to generate RSA-related encryption files, and then fill in the PRIVATE_KEY and PUBLIC_KEY values in the .env file.
Run the code.

Usage in the release environment:

Apply environment variables using commands like export and make sure to set GIN_MODE to "release".
Run the code.

```go
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

```

## Author

👤 tae2089 tae2089002@gmail.com

Github: @tae2089

## 🤝 Contributing

Contributions, issues and feature requests are welcome!
Feel free to check issues page.
