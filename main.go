package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type StockPrice struct {
	CurrentPrice            string
	CurrentChange           string
	CurrentChangePercentage string
}

type ChartResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				HasPrePostMarketData bool    `json:"hasPrePostMarketData"`
				GMTOffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PreviousClose        float64 `json:"previousClose"`
				Scale                int     `json:"scale"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre     TradingPeriod `json:"pre"`
					Regular TradingPeriod `json:"regular"`
					Post    TradingPeriod `json:"post"`
				} `json:"currentTradingPeriod"`
				TradingPeriods  [][]TradingPeriod `json:"tradingPeriods"`
				DataGranularity string            `json:"dataGranularity"`
				Range           string            `json:"range"`
				ValidRanges     []string          `json:"validRanges"`
			} `json:"meta"`
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close []float64 `json:"close"`
					Low   []float64 `json:"low"`
					High  []float64 `json:"high"`
					Open  []float64 `json:"open"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

type TradingPeriod struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	GMToffset int    `json:"gmtoffset"`
}

func fetchPriceFromYahoo(ticker string) (StockPrice, error) {
	var stock StockPrice

	const base_url = "https://query1.finance.yahoo.com/v8/finance/chart/%s?region=US&lang=en-US&includePrePost=false&interval=2m&useYfid=true&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance"
	// const view_mode = "/view/v1"

	// fetch_url := base_url + tickers + view_mode
	fetch_url := fmt.Sprintf(base_url, ticker)

	res, err := http.Get(fetch_url)
	if err != nil {
		return stock, err
	}
	defer res.Body.Close()

	var response ChartResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Chart.Result[0].Meta.ChartPreviousClose)

	// doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	// fmt.Println(doc)

	// if err != nil {
	// 	return stock, err
	// }

	// doc.Find(fmt.Sprintf("[data-field='regularMarketPrice'][data-symbol='%s']", ticker)).Each(func(i int, s *goquery.Selection) {

	// 	stock.CurrentPrice = s.Text()
	// })
	// // doc.Find(fmt.Sprintf("[data-field='regularMarketChange'][data-symbol='%s']", ticker)).Each(func(i int, s *goquery.Selection) {
	// // 	stock.CurrentChange = s.Text()
	// // })
	// doc.Find(fmt.Sprintf("[data-field='regularMarketChangePercent'][data-symbol='%s']", ticker)).Each(func(i int, s *goquery.Selection) {
	// 	stock.CurrentChangePercentage = s.Text()
	// })

	return stock, nil
}

// func getPriceFromYahoo(stock string) string {
// 	var current_price_field = "regularMarketPrice"
// 	var current_change_field = "regularMarketChange"
// 	var current_change_percentage = "regularMarketChangePercent"
// 	return "test"
// }

func main() {
	startTime := time.Now()
	name := flag.String("name", "test", "the name to greet")
	flag.Parse()

	const ticker = "MU"
	_, err := fetchPriceFromYahoo(ticker)
	if err != nil {
		fmt.Printf("Error fetching")
	}

	endTime := time.Now()

	fmt.Printf("Execution time: %s\n", endTime.Sub(startTime))
	if flag.NArg() == 0 {
		fmt.Printf("Hello ,%s\n", *name)
	} else if flag.Arg(0) == "list" {
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.Readdir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Printf("Hello, %s\n", *name)
	}
}
