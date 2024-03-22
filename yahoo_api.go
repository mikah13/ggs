package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



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

func fetchPrice(ticker string) (ChartResponse, error) {
	var response ChartResponse

	const baseUrl = "https://query1.finance.yahoo.com/v8/finance/chart/%s?region=US&lang=en-US&includePrePost=false&interval=2m&useYfid=true&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance"

	fetchUrl := fmt.Sprintf(baseUrl, ticker)

	res, err := http.Get(fetchUrl)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	return response, err

}
