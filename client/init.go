package client

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Model string

var (
	geminiPro       Model = "gemini-pro"
	geminiVisionPro Model = "gemini-pro-vision"
)

func Init() Gemini {
	err := godotenv.Load()
	if err != nil {
		logrus.Error(err)
	}
	key := os.Getenv("API_KEY")
	url := os.Getenv("PROXY_URL")
	c := initClient(key, url)
	return c
}
