package scraper_test

import (
	"testing"

	"github.com/daregod/amazon-co-uk-scraper/scraper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScraper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "scraper")
}

var cases []struct {
	fileName   string
	handlePath string
	result     scraper.AmazonCoUkParsedData
} = []struct {
	fileName   string
	handlePath string
	result     scraper.AmazonCoUkParsedData
}{
	{"1509836071.new_in_stock.html", "/gp/product/1509836071",
		scraper.AmazonCoUkParsedData{Title: "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts", Price: "8.49", Image: "https://images-eu.ssl-images-amazon.com/images/I/61modEZimPL._SX218_BO1,204,203,200_QL40_.jpg", Available: true}},
	{"1787125645.new_and_used.html", "/gp/product/1787125645",
		scraper.AmazonCoUkParsedData{Title: "Go Systems Programming: Master Linux and Unix system level programming with Go", Price: "41.99", Image: "https://images-eu.ssl-images-amazon.com/images/I/41y7-qWywtL._SX218_BO1,204,203,200_QL40_.jpg", Available: true}},
	{"059652692X.used_only.html", "/gp/product/059652692X",
		scraper.AmazonCoUkParsedData{Title: "Asterisk Cookbook", Price: "442.62", Image: "https://images-eu.ssl-images-amazon.com/images/I/51fXTdqAkaL._SX218_BO1,204,203,200_QL40_.jpg", Available: false}},
	{"B000Q646NA.not_a_book.html", "/gp/product/B000Q646NA",
		scraper.AmazonCoUkParsedData{Title: "MAM Start SooTher Suitable 0-2 Months with Sterilisable Travel Case - Pack of 2, Pink", Price: "4.19", Image: "https://images-eu.ssl-images-amazon.com/images/I/41bs5q7xa1L._SY300_QL70_.jpg", Available: true}},
}

func urlsFromCases(baseUrl string) []string {
	urls := make([]string, 0, len(cases))
	for _, cs := range cases {
		urls = append(urls, baseUrl+cs.handlePath)
	}
	return urls
}

func bulkFromCases(baseUrl string) []scraper.AmazonCoUkBulkData {
	results := make([]scraper.AmazonCoUkBulkData, 0, len(cases))
	for _, cs := range cases {
		csRes := cs.result
		results = append(results, scraper.AmazonCoUkBulkData{
			URL:   baseUrl + cs.handlePath,
			Meta:  &csRes,
			Error: nil,
		})
	}
	return results
}
