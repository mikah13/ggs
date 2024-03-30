package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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
		return nil
	case "get-all":
		getWatchlistPrice()
		return nil
	case "list":
		displayWatchlist()
		return nil
	case "add":
		tickers := args[0]
		for _, ticker := range strings.Split(tickers, ",") {
			addTickerToWatchlist(ticker)
		}
		getWatchlistPrice()
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

func addTickerToWatchlist(ticker string) bool {
	var watchList = getWatchList()

	var found = strings.Contains(strings.Join(watchList, ","), ticker)

	if found {
		fmt.Printf("%s is already in the watchlist\n", ticker)
		return false
	}

	watchList = append(watchList, ticker)

	sort.Strings(watchList)

	var fileContent = strings.Join(watchList, ",")

	var operation = updateWatchList(fileContent)

	if !operation {
		fmt.Printf("Something went wrong. %s has not been added to the watchlist !\n", ticker)
		return false
	}

	fmt.Printf("%s has been added to the watchlist !\n", ticker)
	return true
}
