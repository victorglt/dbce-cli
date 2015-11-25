package userapiv0

import (
	"github.com/codegangsta/cli"
)

type Interval struct {
  
}

type QuotesRequest struct {
  Interval, Quantities, osTypes
}


func GetQuotesRequest(c *cli.Context) {

	println("Sending Request....")
	println("Args:", c.Args())
}

func GetQuotesCommand() cli.Command {

	return cli.Command{
		Name:   "getFixedQuotes",
		Usage:  "Get Request for Fixed Contract Quotes",
		Action: GetQuotesRequest,
	}

}
