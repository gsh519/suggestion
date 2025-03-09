package suggestion

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func Suggest() {
	var apiKey = os.Getenv("API_KEY")

	prompt := promptui.Select{
		Label: "どちらを提案してほしいですか？",
		Items: []string{"変数名", "テーブルカラム名"},
	}

	idx, result, err := prompt.Run()

	if err != nil {
		fmt.Println("選択が正しくありません。")
		return
	}

	var desc string
	fmt.Printf("どのような%sを考えてほしいですか？\n", result)
	fmt.Print("→")
	fmt.Scan(&desc)

	var message string
	switch idx {
	case 0:
		message = fmt.Sprintf("下記に適切な%sを5つ程度考えてください。ただし考えた結果以外の文章は不要です。簡潔に答えてください。\n%s", result, desc)
	case 1:
		message = fmt.Sprintf("下記に適切なデータベースの%sを5つ程度考えてください。ただし考えた結果以外の文章は不要です。簡潔に答えてください。\n%s", result, desc)
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	fmt.Print("考え中")

	s := spinner.New(spinner.CharSets[9], 100*time.Microsecond)
	s.Start()

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})

	s.Stop()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(chatCompletion.Choices[0].Message.Content)
}
