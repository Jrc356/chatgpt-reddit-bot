package main

import (
	"fmt"
	"os"

	"github.com/Jrc356/chatgpt-reddit-bot/lib/openai"
	"github.com/Jrc356/chatgpt-reddit-bot/lib/reddit"
)

func main() {
	posts, err := reddit.GetNewPosts(reddit.NewPostsParams{
		Subreddit: "kubernetes",
		N:         1,
		Random:    true,
	})
	if err != nil {
		panic(err)
	}

	openAiKey, ok := os.LookupEnv("OPEN_AI_KEY")
	if !ok {
		panic("OPEN_AI_KEY is unset")
	}

	for _, post := range posts {
		prompt := post.Title + "\n" + post.Content
		fmt.Println("### POST ###\n---")
		fmt.Println(prompt)

		response, err := openai.GenerateText(openai.GenerateTextParams{
			ApiKey:      openAiKey,
			Temperature: 1.0,
			Messages: []openai.Message{
				{
					Role:    openai.User,
					Content: prompt,
				},
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("\n### Response ###\n---")
		fmt.Println(response)
		fmt.Println("END")
	}
}
