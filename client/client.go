package client

import (
	"github.com/imroc/req/v3"
)

var baseUrl = "https://generativelanguage.googleapis.com"

type client struct {
	c   *req.Client
	key string
}

func initClient(key string, proxy string) *client {
	c := req.C().
		SetProxyURL(proxy).
		SetCommonHeader("Content-Type", "application/json")
	return &client{
		c:   c,
		key: key,
	}
}
