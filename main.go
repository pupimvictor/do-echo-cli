package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/pupimvictor/do-echo-cli/client"
	"github.com/pupimvictor/do-echo-cli/client/echo"
	"github.com/pupimvictor/do-echo-cli/models"
	"github.com/urfave/cli"
	"io"
	"log"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
)

type echoClient struct {
	echoer   *client.Echoer
	messages *[]models.Message
	delay    time.Duration
	stdOut   io.Writer
	auth     runtime.ClientAuthInfoWriter
}

func main() {
	err := wrapMain()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func wrapMain() error {
	app := cli.NewApp()
	app.Name = "Echo Client"
	app.Usage = "A client for Echo app"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "echo host server address. format: ip:port",
		},
		cli.StringFlag{
			Name:  "token",
			Usage: "X api token",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "yell",
			Usage:  "yells a message to the echo server",
			Action: runYell,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "msg",
					Usage: "A message to be echoed",
				},
			},
		},
		{
			Name:   "batch",
			Usage:  "yells messages in batch to the echo server",
			Action: runBatch,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "path to a list of messages file",
				},
				cli.Int64Flag{
					Name:  "delay",
					Usage: "delay time in milliseconds between messages",
				},
			},
		},
	}
	return app.Run(os.Args)
}

func runYell(c *cli.Context) error {
	msgStr := c.String("msg")
	msg := &[]models.Message{
		{
			Msg: &msgStr,
		},
	}
	echoCli, err := newEchoClient(c, msg)
	if err != nil {
		return err
	}
	return echoCli.yell()
}

func runBatch(c *cli.Context) error {
	var msgs []models.Message

	filePath := c.String("file")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		msg := scanner.Text()
		msgs = append(msgs, models.Message{Msg: &msg})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	echoCli, err := newEchoClient(c, &msgs)
	if err != nil {
		return err
	}
	return echoCli.yell()
}

func newEchoClient(c *cli.Context, msgs *[]models.Message) (echoClient, error) {
	host := c.GlobalString("host")
	if host == "" {
		host = "127.0.0.1:8000"
	}
	transport := httptransport.New(host, "", nil)
	echoer := client.New(transport, strfmt.Default)

	token := c.GlobalString("token")
	authenticator := httptransport.APIKeyAuth("X-Token", "header", token)

	delay := time.Duration(c.Int64("delay"))

	return echoClient{
		echoer:   echoer,
		auth: authenticator,
		messages: msgs,
		delay:    time.Millisecond * delay,
		stdOut:   os.Stdout,
	}, nil
}

func (e *echoClient) yell() error {
	for _, msg := range *e.messages {
		params := echo.NewEchoParams().WithBody(&msg)
		echoMsg, err := e.echoer.Echo.Echo(params, e.auth)
		if err != nil {
			return errors.New("err calling server: " + err.Error())
		}
		_, err = e.stdOut.Write([]byte(fmt.Sprintf("%s\n", echoMsg.Payload.Echo)))
		if err != nil {
			return errors.New("err writing to stdout:" + err.Error())
		}
		time.Sleep(e.delay)
	}
	return nil
}
