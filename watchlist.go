package main

import (
	"fmt"
	"os"
	"strings"
)

type WatchListOptions struct {
	ConfigPath string
}

// WatchListOption represents a functional option for configuring WatchListOptions.
type WatchListOption func(*WatchListOptions)

// getWatchList retrieves the watch list.
func getWatchList(options ...WatchListOption) []string {
	// Default options
	opt := &WatchListOptions{
		ConfigPath: "./ggs.config",
	}

	// Open the file
	file, err := os.Open(opt.ConfigPath)
	if err != nil {
		// Handle error, e.g., log it or return default value
		return []string{}
	}
	defer file.Close()

	// Read content from the file into a string variable
	var content string
	if _, err := fmt.Fscan(file, &content); err != nil {
		// Handle error, e.g., log it or return default value
		return []string{}
	}

	// Split the content into an array by splitting with ","
	watchList := strings.Split(content, ",")

	// Remove leading and trailing whitespaces from each element
	for i, item := range watchList {
		watchList[i] = strings.TrimSpace(item)
	}

	return watchList
}

func updateWatchList(content string, options ...WatchListOption) bool {
	opt := &WatchListOptions{
		ConfigPath: "./ggs.config",
	}

	// Apply custom options if provided
	for _, option := range options {
		option(opt)
	}

	// Open the file in write mode, create if not exists
	file, err := os.Create(opt.ConfigPath)
	if err != nil {
		// Handle error, e.g., log it
		return false
	}
	defer file.Close()

	// Write content to the file
	_, err = fmt.Fprintf(file, "%s", content)
	if err != nil {
		// Handle error, e.g., log it
		return false
	}

	return true
}
