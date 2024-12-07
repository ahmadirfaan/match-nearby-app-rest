package main

import (
	"os"

	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/ahmadirfaan/match-nearby-app-rest/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}