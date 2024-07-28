package env

import (
	keepalive "awesomo-go/keep-alive"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var envName string = ".env"

func InitEnv() {
	err := godotenv.Load(envName)
	if err != nil {
		keepalive.Start(fmt.Errorf("could not open \"%v\" file", envName))
	}
}

func GetEnv(name string) string {
	return os.Getenv(name)
}
