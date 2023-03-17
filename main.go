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
	api_key := ""
	question := "あなたはプロのソフトウェアエンジニアです。次に示すエラーを解説してください。\n"

    flag.Parse()
    args := flag.Args()

    cli := NewCLI(os.Stdout, os.Stderr)
    _, stderr, _ := cli.Exec(args[0], args[1:]...)

    if stderr != "" {
		question += stderr
		messages = append(messages, Message{
			Role:    "user",
			Content: question,
		})
		response := getResponse(api_key)
		
		fmt.Printf("ChatGPT による解説:")
		fmt.Printf(response.Choices[0].Messages.Content)
	} else {
		fmt.Printf("Command exited with no error.")
	}
    
}
