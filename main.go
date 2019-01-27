package main

import (
	"errors"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/pupimvictor/do-echo-cli/client"
	"github.com/pupimvictor/do-echo-cli/client/echo"
	"github.com/pupimvictor/do-echo-cli/models"
	"github.com/urfave/cli"
	"io"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
)

type echoClient struct {
	echoer   *client.Echoer
	messages *[]models.Message
	delay    time.Duration
	stdOut   io.Writer
}

func newEchoClient(host string, msgs *[]models.Message, delay time.Duration) echoClient {
	transport := httptransport.New(host, "", nil)
	echoer := client.New(transport, strfmt.Default)

	return echoClient{
		echoer: echoer,
		messages: msgs,
		delay: time.Microsecond * delay,
		stdOut: os.Stdout,
	}
}

func main() {
	err := wrapMain()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("")
}

func wrapMain() error {
	app := cli.NewApp()
	app.Name = "Echo Client"
	app.Usage = "A client for Echo app"
	app.Version = "0.0.1"
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
				cli.StringFlag{
					Name:  "host",
					Usage: "echo host server address. format: ip:port",
				},
			},
		},
		{
			Name:   "batch",
			Usage:  "yells a messages in batch",
			Action: runBatch,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "path to a list of messages file",
				},
				cli.IntFlag{
					Name:  "delay",
					Usage: "delay between yells in milliseconds",
				},
				cli.StringFlag{
					Name:  "host",
					Usage: "echo host server address",
				},
			},
		},
	}

	return app.Run(os.Args)
}

func runYell(c *cli.Context) error {
	host := c.String("host")
	if host == "" {
		return errors.New("invalid host address. use do-echo-cli yell --help for help")
	}

	msgStr := c.String("msg")
	msg := &[]models.Message{
		{
			Msg: &msgStr,
		},
	}

	echoCli := newEchoClient(host, msg, 0)
	return echoCli.yell()
}

func (e *echoClient) yell() (error) {
	for _, msg := range *e.messages {
		params := echo.NewEchoParams().WithBody(&msg)
		echoMsg, err := e.echoer.Echo.Echo(params)
		if err != nil {
			return err
		}
		_, err = e.stdOut.Write([]byte(echoMsg.Payload.Echo))
		if err != nil {
			return err
		}
		time.Sleep(e.delay)
	}
	return nil
}

func runBatch(c *cli.Context) error {
	//delay := c.Int("delay")
	return nil
}
