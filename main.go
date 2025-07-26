package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	base_url := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", base_url)

	html, err := getHTML(base_url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(html)
}
