package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

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

	fmt.Printf("Execution time: %s\n", endTime.Sub(startTime))

}
