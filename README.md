### A Simple GO SDK of google gemini

#### Use

1. add ```.env``` file in the root dir (like this)

    ```yaml
   API_KEY=your_apikey
    PROXY_URL=http://127.0.0.1:7890
   ```

2. Simple Get Text Reply

    ```go
    package main
    
    import (
        "context"
        "fmt"
        "gemini-go/client"
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

3. Stream Mode


#### Implement

- [x] Common Msg Reply
- [x] Stream Msg Reply
- [ ] Image Vision

#### Notice

change vpn area to USA
