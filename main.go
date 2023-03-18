package main

import (
    "bytes"
    "context"
    "flag"
    "fmt"
    "io"
    "os"
    "os/exec"
    "strings"
	"log"

	"github.com/joho/godotenv"
	"github.com/jeandeaual/go-locale"
	"github.com/emvi/iso-639-1"
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
			fmt.Printf("Error: Please specify the target command.\nUsage: descit <your command>")
		}
	}()


	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
        return
	}

	err = godotenv.Load(home + "/.config/descit/.env")
	if err != nil {
		log.Fatal("Error: Couldn't retrieve OpenAI's token. Please set the token in `~/.config/descit/.env`.\nREADME: https://github.com/yude/descit")
	}

	userLocales, err := locale.GetLocales()

	api_key := os.Getenv("TOKEN")
	question := "You are a full-time senior software engineer. Describe the following error in " + iso6391.Name(userLocales[len(userLocales) - 1]) + "\n"

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
		fmt.Printf("Command exited with non-zero code. Calling ChatGPT sensei...")
		response := getResponse(api_key)
		fmt.Printf("Explained by ChatGPT:")
		fmt.Printf(response.Choices[0].Messages.Content)
	} else {
		fmt.Printf("Command exited with no error.")
	}
    
}
