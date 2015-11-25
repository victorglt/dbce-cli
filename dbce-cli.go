package main

import (
	"github.com/codegangsta/cli"
	"github.com/victorglt/dbce-cli/userapiv0"
	"os"
)

const Version = "0.1"

func main() {
	app := cli.NewApp()

	app.Name = "dbce-cli"
	app.Author = "Deutsche Börse Cloud Exchange (victor.galante@cloud.exchange)"
	app.Version = Version
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		userapiv0.GetQuotesCommand(),
	}

	app.Run(os.Args)
}
