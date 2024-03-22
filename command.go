package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	usage = `Specify a command to execute:
  - get: search stock price
  - list: display your watchlist
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
	for _, ticker := range tickersArray {
		res, err := fetchPrice(ticker)
		if err != nil {
			fmt.Printf("Error fetching")
		}
	}
}

func getWatchlistPrice() {
	var watchList = getWatchList()
	for _, ticker := range watchList {
		res, err := fetchPrice(ticker)
		if err != nil {
			fmt.Printf("Error fetching")
		}
	}
}

func displayWatchlist() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
