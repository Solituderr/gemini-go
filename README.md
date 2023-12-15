### A Simple GO SDK of google gemini

#### Use

Simple Get Text Reply

```go
package main

import (
	"context"
	"fmt"
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
	s := tx.Chat(contents, ctx)
	fmt.Println(s)
	cancel()
}

```

Stream Mode


#### Implement

- [x] Common Msg Reply
- [x] Stream Msg Reply
- [ ] Image Vision
