package scraper

import (
	"bytes"

	"github.com/go-playground/log"
	"github.com/gocolly/colly"
)

// AmazonCoUkBulkData contain AmazonCoUkParsedData with meta information
type AmazonCoUkBulkData struct {
	URL   string                `json:"url"`
	Meta  *AmazonCoUkParsedData `json:"meta,omitempty"`
	Error *string               `json:"error,omitempty"`
}

// Processor envelope
type Processor interface {
	ProcessUrls(urls []string) []AmazonCoUkBulkData
}

// AmazonCoUkProcessor is base Processor implementation
type AmazonCoUkProcessor struct {
	// Exported to allow flexibility tune
	Collector *colly.Collector
}

// NewProcessor construct new Processor with default (limited) collector
func NewProcessor() Processor {
	return AmazonCoUkProcessor{
		Collector: colly.NewCollector(
			colly.AllowedDomains("amazon.co.uk", "www.amazon.co.uk"),
		),
	}
}

// ProcessUrls will run parser on each url
func (pr AmazonCoUkProcessor) ProcessUrls(urls []string) []AmazonCoUkBulkData {
	result := make([]AmazonCoUkBulkData, 0, len(urls))

	pr.Collector.UserAgent = "Ad-Blocker/1 CFNetwork/758.5.3 Darwin/15.6.0"

	pr.Collector.OnError(func(r *colly.Response, err error) {
		errSt := err.Error()
		result = append(result, AmazonCoUkBulkData{
			URL:   r.Request.URL.String(),
			Error: &errSt,
		})
	})

	pr.Collector.OnResponse(func(r *colly.Response) {
		item := AmazonCoUkBulkData{
			URL: r.Request.URL.String(),
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
		err := pr.Collector.Visit(url)
		if err != nil {
			log.WithError(err).WithField("url", url).Alert("Cannot visit url while processing") //nolint: lll
		}
	}
	pr.Collector.Wait()
	return result
}
