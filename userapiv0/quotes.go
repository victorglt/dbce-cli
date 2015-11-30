package quotes

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
)

const (
	url          = "https://api.cloud.exchange/v0"
	getQuotesUrl = "/get-fixed-quotes"
)

func LogError(err error) {
	if err != nil {
		println(err)
	}
}

func GetQuotesRequest(c *cli.Context) {

	println("Sending Request....")

	resp, err := http.Get(url + getQuotesUrl)

	LogError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	LogError(err)

	println(string(body))

}

func GetQuotesCommand() cli.Command {

	return cli.Command{
		Name:   "getFixedQuotes",
		Usage:  "Get Request for Fixed Contract Quotes",
		Action: GetQuotesRequest,
	}

}
