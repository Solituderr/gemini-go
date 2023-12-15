package main

import (
	"context"
	"gemini/client"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	tx := client.Init().GetTextFunc()
	contents := []client.Content{
		{
			Role:  "user",
			Parts: []client.Text{{"帮我写一段故事"}},
		},
	}
	tx.ChatStream(contents, ctx)
	cancel()
}
