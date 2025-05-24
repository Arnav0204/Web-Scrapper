package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func scrape(c *colly.Collector, url string) string {
	var content string

	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println("Tag:", e.Name, "Text:", e.Text)
		content += e.Text
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
