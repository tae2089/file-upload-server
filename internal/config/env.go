package config

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	projectDir := getProjectDir()
	fmt.Println("loading.env file")
	err := godotenv.Load(projectDir + "/.env")
	if err != nil {
		log.Panic("Error loading.env file")
	}
}

func getProjectDir() string {
	projectDir := ""
	_, filename, _, _ := runtime.Caller(0)
	projectDir = path.Join(path.Dir(filename), "../..")
	return projectDir
}
