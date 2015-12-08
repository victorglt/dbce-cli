package main

import (
	"github.com/codegangsta/cli"
	"github.com/victorglt/dbce-cli/configuration"
	"github.com/victorglt/dbce-cli/userapiv0"
	"os"
)

const Version = "0.1"

func main() {
	configuration.SetupConfig()

	app := cli.NewApp()

	app.Name = "dbce-cli"
	app.Author = "Deutsche Börse Cloud Exchange (victor.galante@cloud.exchange)"
	app.Version = Version
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		quotes.GetQuotesCommand(),
	}

	app.Run(os.Args)
}
