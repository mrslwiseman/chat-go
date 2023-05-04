`chat-go "write a readme markdown file for this script: $(cat chat-go.go)"`

---

# Chat-bot using OpenAI API

This is a command-line interface chat-bot built with the OpenAI API. It uses the GPT-3.5 turbo model to generate responses based on the input given by the user.

## Installation

1. Clone the repository
2. Make sure you have Go installed and properly set up in your system.
3. Run `go mod tidy` to install dependencies
4. Set up an OpenAI API key and store it in an environmental variable named "OPENAI_API_KEY".

## Usage

To run the program, navigate to the directory where the script is saved and enter the following command:

`go run main.go`

This will start the chat-bot interface. Simply type in your message and hit enter. The bot will generate a response based on your input.

To exit the chat-bot interface, type "exit" or press "\q" and hit enter.

You can also use the CLI non-interactively by passing the message as an argument when running the program. For example:

`go run main.go "Hello, how are you?"`

This will generate a response from the bot based on the message "Hello, how are you?".

You can also use the output of the program in a bash script by piping it to a file, like so:

`go run main.go "Hello, how are you?" > output.txt`

This will save the output in a file named "output.txt".
