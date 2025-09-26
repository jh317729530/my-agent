package main

import (
	"context"
	"fmt"
	"io"

	"my-agent/llm"

	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()
	messages := []*schema.Message{
		{
			Role:    schema.System,
			Content: "You are a helpful AI assistant. Be concise in your responses.",
		},
		{
			Role:    schema.User,
			Content: "What is the capital of France?",
		},
	}
	resp, err := llm.Model.Stream(ctx, messages)
	if err != nil {
		panic(err)
	}

	i := 0
	for {
		message, err := resp.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("message[%d]: %+v", i, message.Content))
		i++
	}
}
