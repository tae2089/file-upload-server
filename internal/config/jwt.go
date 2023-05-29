package config

import (
	"os"

	"github.com/tae2089/file-upload-server/domain"
)

var (
	jwtKey domain.JwtKey
)

func GetJwtKey() domain.JwtKey {
	if jwtKey.PrivateKey == "" || jwtKey.PublicKey == "" {
		jwtKey.PrivateKey = os.Getenv("PRIVATE_KEY")
		jwtKey.PublicKey = os.Getenv("PUBLIC_KEY")
	}
	return jwtKey
}
