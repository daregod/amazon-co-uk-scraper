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

type Processor interface {
	ProcessUrls(urls []string) []AmazonCoUkBulkData
}

type AmazonCoUkProcessor struct {
	Collector *colly.Collector
}

func NewProcessor() Processor {
	return AmazonCoUkProcessor{
		Collector: colly.NewCollector(
			colly.AllowedDomains("amazon.co.uk", "www.amazon.co.uk"),
		),
	}
}

func (pr AmazonCoUkProcessor) ProcessUrls(urls []string) []AmazonCoUkBulkData {
	result := make([]AmazonCoUkBulkData, 0, len(urls))

	pr.Collector.UserAgent = "Ad-Blocker/1 CFNetwork/758.5.3 Darwin/15.6.0"

	pr.Collector.OnError(func(r *colly.Response, err error) {
		errSt := err.Error()
		result = append(result, AmazonCoUkBulkData{
			Url:   r.Request.URL.String(),
			Error: &errSt,
		})
	})

	pr.Collector.OnResponse(func(r *colly.Response) {
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
		pr.Collector.Visit(url)
	}
	pr.Collector.Wait()
	return result
}
