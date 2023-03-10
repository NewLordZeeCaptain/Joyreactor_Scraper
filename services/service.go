package services

import (
	"github.com/gocolly/colly"
)

// Basic Scraping Function. Returning Status code
func Scraper(URL string, content map[string]string) int {
	c := colly.NewCollector(
		colly.AllowedDomains(URL),
	)

	// c.OnRequest(func(r *colly.Request) {})
	var StatusCode int
	c.OnResponse(func(r *colly.Response) {
		StatusCode = r.StatusCode
	})

	c.OnHTML("div.image a[href] img", func(h *colly.HTMLElement) {
		link := h.Attr("src")

		title := h.Attr("title")
		content[title] = link
		// fmt.Printf("Title: %s -> Link: %s\n", title, link)

	})

	c.Visit("https://" + URL)
	return StatusCode
}
