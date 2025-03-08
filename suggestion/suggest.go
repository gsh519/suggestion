package suggestion

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func Suggest() {
	var apiKey = os.Getenv("API_KEY")

	var variableType string
	fmt.Print("どちらを提案してほしいですか？変数名？テーブルカラム名？：")
	fmt.Scan(&variableType)

	switch variableType {
	case "変数名":
		fmt.Println("あなたが選んだのは変数名ですね")
	case "テーブルカラム名":
		fmt.Println("あなたが選んだのはテーブルカラム名ですね")
	}

	var desc string
	fmt.Printf("どのような%sを考えてほしいですか？\n", variableType)
	fmt.Print("→ ")
	fmt.Scan(&desc)

	var message string
	switch variableType {
	case "変数名":
		message = fmt.Sprintf("下記に適切な%sを5つ程度考えてください。ただし考えた結果以外の文章は不要です。簡潔に答えてください。\n%s", variableType, desc)
	case "テーブルカラム名":
		message = fmt.Sprintf("下記に適切なデータベースの%sを5つ程度考えてください。ただし考えた結果以外の文章は不要です。簡潔に答えてください。\n%s", variableType, desc)
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	fmt.Println("考え中...")

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(chatCompletion.Choices[0].Message.Content)
}
