package scraper

import (
	"bytes"

	"github.com/gocolly/colly"
)

type AmazonCoUkBulkData struct {
	Url   string                `json:"url"`
	Meta  *AmazonCoUkParsedData `json:"meta,omitempty"`
	Error *string               `json:"error,omitempty"`
}

func ProcessUrls(urls []string) []AmazonCoUkBulkData {
	result := make([]AmazonCoUkBulkData, 0, len(urls))

	c := colly.NewCollector(
		colly.AllowedDomains("amazon.co.uk", "www.amazon.co.uk"),
	)
	c.UserAgent = "Ad-Blocker/1 CFNetwork/758.5.3 Darwin/15.6.0"

	c.OnError(func(r *colly.Response, err error) {
		errSt := err.Error()
		result = append(result, AmazonCoUkBulkData{
			Url:   r.Request.URL.String(),
			Error: &errSt,
		})
	})

	c.OnResponse(func(r *colly.Response) {
		item := AmazonCoUkBulkData{
			Url: r.Request.URL.String(),
		}
		pd := Parse(bytes.NewBuffer(r.Body))
		err := pd.Check()
		if err == nil {
			item.Meta = &pd
		} else {
			estr := err.Error()
			item.Error = &estr
		}
		result = append(result, item)
	})
	for _, url := range urls {
		c.Visit(url)
	}
	c.Wait()
	return result
}
