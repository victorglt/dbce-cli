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
	"strings"
	"time"
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
	Start string `json:"start, omitempty"`
	End   string `json:"end, omitempty"`
}

type Quantities struct {
	Compute string `json:"compute, omitempty"`
	Storage string `json:"storage, omitempty"`
	Memory  string `json:"memory, omitempty"`
}

type FixedQuoteRequest struct {
	Quantities *Quantities `json:"quantities, omitempty"`
	Interval   *Interval   `json:"interval, omitempty"`
}

type FixedQuote struct {
	Id         string `json:"id, omitempty"`
	TotalPrice string `json:"totalPrice, omitempty"`
}

type Response struct {
	Data []FixedQuote `json:"data, omitempty"`
}

func ReadQuantities() *Quantities {
	quant := new(Quantities)
	print("dbce-cli>> Type compute qty (default: 1): ")
	fmt.Scanf("%s", &quant.Compute)
	if strings.Compare(quant.Compute, "") == 0 {
		quant.Compute = "1"
	}

	print("dbce-cli>> Type storage qty (default: 1): ")
	fmt.Scanf("%s", &quant.Storage)
	if strings.Compare(quant.Storage, "") == 0 {
		quant.Storage = "1"
	}

	print("dbce-cli>> Type memory qty (default: 1): ")
	fmt.Scanf("%s", &quant.Memory)
	if strings.Compare(quant.Memory, "") == 0 {
		quant.Memory = "1"
	}

	return quant
}

func ReadInterval() *Interval {
	interval := new(Interval)
	print("dbce-cli>> Type start date: ")
	fmt.Scanf("%s", &interval.Start)
	print("dbce-cli>> Type end date: ")
	fmt.Scanf("%s", &interval.End)
	return interval
}

func GetDefaultInterval() *Interval {
	interval := new(Interval)

	now := time.Now()

	interval.Start = now.Format(time.RFC3339)
	interval.End = now.Format(time.RFC3339)

	return interval
}

func GetQuotesRequest(c *cli.Context) {

	quant := ReadQuantities()

	var interval *Interval

	print("dbce-cli>> Specify Interval [Default: yesterday-today] ? (Y/n) ")
	var yn string
	fmt.Scanf("%s", &yn)

	if yn == "Y" || yn == "y" {
		interval = ReadInterval()
	} else {
		interval = GetDefaultInterval()
	}

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
