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

