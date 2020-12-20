package crawler

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gocolly/colly"
	"github.com/olivere/elastic"
)

var (
	searchString  string         = "https://www.goodreads.com/quotes/"
	contentRegexp *regexp.Regexp = regexp.MustCompile("“(.+?)”")
)

// Quote type
type Quote struct {
	Content string                `json:"content"`
	Author  string                `json:"author"`
	Created time.Time             `json:"created,omitempty"`
	Suggest *elastic.SuggestField `json:"suggest_field,omitempty"`
}

// Run crawler
func Run(amount int) []Quote {
	var quotes []Quote

	c := colly.NewCollector(
		colly.AllowedDomains("www.goodreads.com"),
	)

	c.OnHTML(".quoteDetails", func(e *colly.HTMLElement) {
		if len(quotes) >= amount {
			return
		}
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
			Created: time.Now(),
		})
	})

	c.OnHTML(".next_page", func(e *colly.HTMLElement) {
		if len(quotes) < amount {
			e.Request.Visit(e.Attr("href"))
		}
	})

	fmt.Println("Launching Crawler...")
	c.Visit(searchString)
	return quotes
}
