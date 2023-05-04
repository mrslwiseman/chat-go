package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

var KEY string = os.Getenv("OPENAI_API_KEY")

var client = openai.NewClient(KEY)

func prompt() string {
	fmt.Printf(">>> ")
	rdr := bufio.NewReader(os.Stdin)
	input, err := rdr.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return input
}

var convo = []openai.ChatCompletionMessage{}

func createChat(convo []openai.ChatCompletionMessage) openai.ChatCompletionResponse {

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: convo,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		panic(err)
	}

	return resp
}

func main() {
	// This allows the CLI to not be interactive so you can pipe the output eg $ chat-go write a bash loop > loop.sh
	if len(os.Args) > 1 {
		resp := createChat([]openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: strings.Join(os.Args[1:], " ")}})

		fmt.Println(resp.Choices[0].Message.Content)
		os.Exit(0)
	}

	for {
		userInput := prompt()

		if userInput == "exit\n" || userInput == "\\q\n" {
			os.Exit(0)
		}

		convo = append(convo, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: userInput})

		resp := createChat(convo)

		convo = append(convo, resp.Choices[0].Message)

		fmt.Println("------------------------------------------------------")
		fmt.Println(resp.Choices[0].Message.Content)
		fmt.Println("------------------------------------------------------")
	}

}
