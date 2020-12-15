package crawler

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

var (
	searchString  string         = "https://www.goodreads.com/quotes/"
	contentRegexp *regexp.Regexp = regexp.MustCompile("“(.+?)”")
)

// Quote type
type Quote struct {
	Content string
	Author  string
}

// Run crawler
func Run(amount int) []Quote {
	var quotes []Quote

	c := colly.NewCollector(
		colly.AllowedDomains("www.goodreads.com"),
	)

	c.OnHTML(".quoteDetails", func(e *colly.HTMLElement) {
		res := contentRegexp.FindAllStringSubmatch(e.ChildText("div.quoteText"), -1)
		if len(res) < 1 {
			return
		}
		if len(res[0]) < 1 {
			return
		}
		quotes = append(quotes, Quote{
			Content: res[0][0],
			Author:  e.ChildText(".authorOrTitle"),
		})
	})

	c.OnHTML(".next_page", func(e *colly.HTMLElement) {
		if len(quotes) < amount {
			e.Request.Visit(e.Attr("href"))
		}
	})

	fmt.Println("Launching Scraper !")
	c.Visit(searchString)
	return quotes
}
