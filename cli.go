package echocli

import (
	"fmt"
	"github.com/pupimvictor/do-echo-cli/client"
	"github.com/pupimvictor/do-echo-cli/models"
	"github.com/urfave/cli"
	"os"
)

type echoclient struct {
	echoer client.Echoer
	messages []*models.Message
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
	app.Commands = []cli.Command{
		{
			Name:   "yell",
			Usage:  "yells a message to the echo server",
			Action: runYell,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "msg",
					Usage: "A message to be echoed",
				},
				cli.StringFlag{
					Name: "host",
					Usage: "echo host server address",
				},
			},
		},
		{
			Name:   "batch",
			Usage:  "yells a messages in batch",
			Action: runBatch,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "file",
					Usage: "path to a list of messages file",
				},
				cli.StringFlag{
					Name: "host",
					Usage: "echo host server address",
				},
			},
		},

	}
	return nil
}

func runYell(c *cli.Context) error {
	return nil
}

func runBatch(c *cli.Context) error {
	return nil
}
