package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	name := flag.String("name", "test", "the name to greet")
	flag.Parse()

	var watchList = getWatchList()
	for _, ticker := range watchList {
		res, err := fetchPrice(ticker)
		if err != nil {
			fmt.Printf("Error fetching")
		}
		fmt.Println(ticker, ":", res.Chart.Result[0].Meta.RegularMarketPrice)

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
