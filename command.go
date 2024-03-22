package main

import (
	"fmt"
	"os"

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
		fmt.Println("search repos command", args[0])
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

func displayWatchlist() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}