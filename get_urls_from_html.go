package main

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	urls := []string{}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if strings.Contains(a.Val, "http") {
						urls = append(urls, a.Val)
						break
					} else {
						normalizedURL, err := normalizeURL(a.Val)
						if err != nil {
							return nil, err
						}
						formattedURL := rawBaseURL + normalizedURL
						urls = append(urls, formattedURL)
						break
					}
				}
			}
		}
	}

	return urls, nil
}
