package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	content := make(map[string]string)
	URL := "joyreactor.cc"
	Scraper(URL, content)
	CsvWriter(content)
}

func CsvWriter(content map[string]string) {
	recordFile, err := os.Create("./records.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(recordFile)

	writer.Write([]string{"title", "link"})
	for k, v := range content {
		writer.Write([]string{k, v})
	}
	writer.Flush()

}

func CvsRead() {
	recordFile, err := os.Open("./records.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(recordFile)
	record, _ := reader.ReadAll()
	fmt.Println(record)
}

type Item struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func Scraper(URL string, content map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains(URL),
	)
	c.Post(URL+"/login", map[string]string{"username": "zeecape", "password": "kirkirkirq11111"})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visit", r.URL)

	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnHTML("div.image a[href] img", func(h *colly.HTMLElement) {
		link := h.Attr("src")

		title := h.Attr("title")
		content[title] = link
		// fmt.Printf("Title: %s -> Link: %s\n", title, link)

	})

	c.Visit("https://" + URL)

}
