package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	usage = `Specify a command to execute:
  - get: display stock price in a table
  - get-all: display watchlist in a table
  - list: editable list of all tickers in watchlist
  - add: add ticker to watchlist`
)

func executeCommand(command string, args []string) error {
	switch command {
	case "get":
		tickers := args[0]
		getTickersPrice(tickers)
		fmt.Println("search repos command", args[0])
		return nil

	case "get-all":
		getWatchlistPrice()
		return nil
	case "list":
		displayWatchlist()
		return nil
	case "add":
		fmt.Println("search users command")
		return nil
	default:
		return fmt.Errorf("invalid command: '%s'\n\n%s\n", command, usage)
	}
}

func getTickersPrice(tickers string) {
	tickersArray := strings.Split(tickers, ",")
	var stocks []ChartResponse
	for _, ticker := range tickersArray {
		res, err := fetchPrice(ticker)
		if err != nil {
			fmt.Printf("Error fetching")
		}
		stocks = append(stocks, res)
	}
	getTable(stocks)
}

func getWatchlistPrice() {
	var watchList = getWatchList()
	var stocks []ChartResponse
	for _, ticker := range watchList {
		res, err := fetchPrice(ticker)
		if err != nil {
			fmt.Printf("Error fetching")
		}
		stocks = append(stocks, res)
	}
	getTable(stocks)
}

func displayWatchlist() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func getTable(stocks []ChartResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Ticker", "Last Price", "Change", "Change %", "Previous Close", "Currency"})

	for _, stock := range stocks {
		row := getRow(stock)
		t.AppendRow(row)
	}
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.Render()
}

func getRow(stock ChartResponse) table.Row {
	data := stock.Chart.Result[0].Meta
	diff := data.RegularMarketPrice - data.PreviousClose
	ticker := data.Symbol
	lastPrice := data.RegularMarketPrice
	change := appendPlus(diff)
	changePercent := appendPlus(diff / data.PreviousClose * 100)
	currency := data.Currency
	previousClose := data.PreviousClose

	return table.Row{ticker, lastPrice, change, changePercent, previousClose, currency}
}

func appendPlus(num float64) string {
	if num >= 0 {
		return fmt.Sprintf("+%.2f", num)
	}
	return fmt.Sprintf("%.2f", num)
}
