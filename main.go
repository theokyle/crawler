package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("no website provided")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}

	base_url := os.Args[1]
	maxConcurrency := 5

	cfg, err := configure(base_url, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", base_url)

	cfg.wg.Add(1)
	go cfg.crawlPage(base_url)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
