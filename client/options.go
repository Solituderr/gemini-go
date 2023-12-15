package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

type Gemini interface {
	GeminiPro
	GeminiVisionPro
}

func (c *client) GetTextFunc() GeminiPro {
	c.c.SetBaseURL("https://generativelanguage.googleapis.com")
	return GeminiPro(c)
}

func (c *client) GetImageFunc() GeminiVisionPro {
	c.c.SetBaseURL("https://generativelanguage.googleapis.com/v1beta/models/gemini-pro-vision")
	return GeminiVisionPro(c)
}

type GeminiPro interface {
	GetTextFunc() GeminiPro
	SetLimit(m map[string]string) GeminiPro
	Chat(contents []Content, ctx context.Context) string
	ChatStream(contents []Content, ctx context.Context) *bufio.Scanner
}

type GeminiVisionPro interface {
	GetImageFunc() GeminiVisionPro
}

func (c *client) SetLimit(m map[string]string) GeminiPro {
	c.c.SetCommonFormData(m)
	return c
}

func (c *client) Chat(contents []Content, ctx context.Context) string {
	url := "/v1beta/models/gemini-pro:generateContent"
	msg := Messages{Contents: contents}
	resp := c.c.Post(url).SetQueryParam("key", c.key).SetBody(msg).Do(ctx)
	data := resp.String()
	s := gjson.Get(data, "candidates.0.content.parts.0.text").String()
	return s
}

func (c *client) ChatStream(contents []Content, ctx context.Context) *bufio.Scanner {
	url := "/v1beta/models/gemini-pro:streamGenerateContent"
	msg := Messages{Contents: contents}
	resp := c.c.DisableAutoReadResponse().Post(url).SetQueryParam("key", c.key).SetBody(msg).Do(ctx)
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "text") {
			fmt.Println(line)
		}
	}
	return scanner
}
