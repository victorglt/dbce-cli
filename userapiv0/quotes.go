package quotes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/victorglt/dbce-cli/configuration"
	"io/ioutil"
	"log"
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
		log.Println(err)
	}
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
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
	Id         string `json:"id"`
	TotalPrice string `json:"totalPrice"`
}

type Response struct {
	Data []FixedQuote `json:"data"`
}

func GetQuotesRequest(c *cli.Context) {

	quant := Quantities{}
	print("dbce-cli>> Type compute qty (empty): ")
	fmt.Scanf("%s", &quant.Compute)
	print("dbce-cli>> Type storage qty (empty): ")
	fmt.Scanf("%s", &quant.Storage)

	interval := Interval{}
	print("dbce-cli>> Type start date (yesterday): ")
	fmt.Scanf("%s", &interval.Start)
	print("dbce-cli>> Type end date (today): ")
	fmt.Scanf("%s", &interval.End)

	filter := &FixedQuoteRequest{
		Quantities: quant,
		Interval:   interval,
	}
	jsonFilter, err := json.Marshal(&filter)

	println(string(jsonFilter))

	req, err := http.NewRequest("POST", configuration.Context.Url+getQuotesUrl, bytes.NewBuffer(jsonFilter))
	req.Header.Add("DBCE-ApiKey", configuration.Context.Key)

	resp, err := client.Do(req)

	LogFatal(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	//quotes, err := json.Unmarshal(body, Response)
	LogFatal(err)

	println(string(body))

}

func GetQuotesCommand() cli.Command {

	return cli.Command{
		Name:   "getFixedQuotes",
		Usage:  "Get Request for Fixed Contract Quotes",
		Action: GetQuotesRequest,
	}

}
