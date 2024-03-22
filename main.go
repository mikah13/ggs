package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	startTime := time.Now()
	flag.Parse()

	if len(flag.Args()) < 1 {
		os.Exit(1)
	}

	command := flag.Args()[0]

	err := executeCommand(command, flag.Args()[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	endTime := time.Now()

	// startTime := time.Now()
	// name := flag.String("name", "test", "the name to greet")
	// flag.Parse()

	// var watchList = getWatchList()

	// p := tea.NewProgram(initialModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }

	// for _, ticker := range watchList {
	// 	res, err := fetchPrice(ticker)
	// 	if err != nil {
	// 		fmt.Printf("Error fetching")
	// 	}
	// 	fmt.Println(ticker, ":", res.Chart.Result[0].Meta.RegularMarketPrice)

	// }

	// endTime := time.Now()

	fmt.Printf("Execution time: %s\n", endTime.Sub(startTime))
	// if flag.NArg() == 0 {
	// 	fmt.Printf("Hello ,%s\n", *name)
	// } else if flag.Arg(0) == "list" {
	// 	files, _ := os.Open(".")
	// 	defer files.Close()

	// 	fileInfo, _ := files.Readdir(-1)
	// 	for _, file := range fileInfo {
	// 		fmt.Println(file.Name())
	// 	}
	// } else {
	// 	fmt.Printf("Hello, %s\n", *name)
	// }
}
