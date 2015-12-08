package quotes

import (
	"bytes"
	"encoding/json"
	"github.com/codegangsta/cli"
	"github.com/victorglt/dbce-cli/configuration"
	"io/ioutil"
	"net/http"
)

const (
	getQuotesUrl = "/get-fixed-quotes"
)

var (
	client = &http.Client{}
)

func LogError(err error) {
	if err != nil {
		println(err)
	}
}

type Interval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Quantities struct {
	Compute string `json:"compute"`
	Storage string `json:"storage"`
}
type FixedQuoteRequest struct {
	Quantities Quantities `json:"quantities"`
	Interval   Interval   `json:"interval"`
}

type FixedQuote struct {
	Id         String `json:"id"`
	TotalPrice string `json:"totalPrice"`
}

type Response struct {
	Data []FixedQuote `json:"data"`
}

func GetQuotesRequest(c *cli.Context) {

	println(configuration.Context.Url + getQuotesUrl)

	filter := &FixedQuoteRequest{
		Quantities: Quantities{Compute: "10", Storage: "10"},
		Interval:   Interval{Start: "2015-12-08T23:00:00.000Z", End: "2015-12-09T23:00:00.000Z"},
	}
	jsonFilter, err := json.Marshal(&filter)

	println(string(jsonFilter))

	req, err := http.NewRequest("POST", configuration.Context.Url+getQuotesUrl, bytes.NewBuffer(jsonFilter))
	req.Header.Add("DBCE-ApiKey", configuration.Context.Key)

	resp, err := client.Do(req)

	LogError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	//quotes, err := json.Unmarshal(body, Response)
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
