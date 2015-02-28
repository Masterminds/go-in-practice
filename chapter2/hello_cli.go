package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}
	app.Action = func(c *cli.Context) {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)
	}

	app.Run(os.Args)
}
