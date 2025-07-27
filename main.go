package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("no website provided")
		return
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	base_url := os.Args[1]
	maxConcurrency, err := strconv.ParseInt(os.Args[2], 10, 0)
	if err != nil {
		fmt.Printf("Error - parsing maxConcurrency %v", err)
	}
	maxPages, err := strconv.ParseInt(os.Args[3], 10, 0)
	if err != nil {
		fmt.Printf("Error - parsing maxPages %v", err)
	}

	cfg, err := configure(base_url, int(maxConcurrency), int(maxPages))
	if err != nil {
		fmt.Printf("Error - configure %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", base_url)

	cfg.wg.Add(1)
	go cfg.crawlPage(base_url)
	cfg.wg.Wait()

	printReport(cfg.pages, base_url)
}
