package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/jeandeaual/go-locale"
	"github.com/joho/godotenv"
)

type CLI struct {
	out io.Writer
	err io.Writer
}

func NewCLI(out, err io.Writer) *CLI {
	return &CLI{
		out: out,
		err: err,
	}
}

func (c *CLI) Exec(cmd string, args ...string) (stdout string, stdin string, err error) {
	cmdctx := exec.CommandContext(context.TODO(), cmd, args...)

	var outbuf, errbuf bytes.Buffer
	cmdctx.Stdout = io.MultiWriter(c.out, &outbuf)
	cmdctx.Stderr = io.MultiWriter(c.err, &errbuf)

	if err := cmdctx.Run(); err != nil {
		return outbuf.String(), errbuf.String(), fmt.Errorf("executing %s %s: %w", cmd, strings.Join(args, ""), err)
	}
	return outbuf.String(), errbuf.String(), nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("⚠️ Error: Please specify the target command.\n  Usage: descit <your_command>")
		}
	}()

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = godotenv.Load(home + "/.config/descit/.env")
	if err != nil {
		log.Fatal("⚠️ Error: Couldn't retrieve OpenAI's token. Please set the token in `~/.config/descit/.env`.\nREADME: https://github.com/yude/descit")
	}

	userLocales, err := locale.GetLocales()

	api_key := os.Getenv("TOKEN")
	question := "あなたはプロのソフトウェアエンジニアです。以下に続くエラーを分かりやすく説明してください。ただし、この説明は、自然言語の一つである" + userLocales[len(userLocales)-1] + "を使用して行ってください。\n"

	flag.Parse()
	args := flag.Args()

	cli := NewCLI(os.Stdout, os.Stderr)
	_, stderr, err := cli.Exec(args[0], args[1:]...)

	if stderr != "" {
		question += stderr
		messages = append(messages, Message{
			Role:    "user",
			Content: question,
		})
		fmt.Println("😩 Command exited with non-zero code. Calling ChatGPT sensei...")
		response := getResponse(api_key)
		fmt.Println("💬 Explanation by ChatGPT (Model: " + response.Model + ", Language: " + userLocales[len(userLocales)-1] + "):")
		fmt.Println(strings.Replace(response.Choices[0].Messages.Content, "\n\n", "", -1))
	} else {
		fmt.Printf("✅ Command exited with no error.")
	}

}
