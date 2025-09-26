package llm

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/components/model"
)

var (
	Model model.ToolCallingChatModel
)

func init() {
	var err error
	Model, err = deepseek.NewChatModel(context.Background(), &deepseek.ChatModelConfig{
		APIKey: os.Getenv("DEEPSEEK_API_KEY"),
		Model:  "deepseek-chat",
	})
	if err != nil {
		panic(err)
	}
}
