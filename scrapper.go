package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func scrape(c *colly.Collector, url string) string {
	var content string

	c.OnHTML("h1, h2, h3, h4, h5, h6, p, li", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		if text != "" {
			content += text + "\n"
		}
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	return content
}
