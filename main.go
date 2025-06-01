package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Welcome To Web Scrapper")
	fmt.Println("Enter the start URL")
	var url string
	fmt.Scanln(&url)
	c := colly.NewCollector()
	beginScrapping(c, url)
}

func beginScrapping(c *colly.Collector, url string) {
	//var urlList []string
	var content string
	//urlList = append(urlList, url)
	content = scrape(c, url)
	fmt.Println(content)
	fileWriter(content)

}

func fileWriter(content string) {
	f, err := os.Create("./file.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

}
