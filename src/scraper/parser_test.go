package scraper_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Parser", func() {
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
	Describe("Parse", func() {
		cases := []struct {
			filename string
			result   scraper.AmazonCoUkParsedData
		}{
			{"1509836071.new_in_stock.html",
				scraper.AmazonCoUkParsedData{Title: "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts", Price: "8.49", Image: "https://images-eu.ssl-images-amazon.com/images/I/61modEZimPL._SX218_BO1,204,203,200_QL40_.jpg", Available: true}},
			{"1787125645.new_and_used.html",
				scraper.AmazonCoUkParsedData{Title: "Go Systems Programming: Master Linux and Unix system level programming with Go", Price: "41.99", Image: "https://images-eu.ssl-images-amazon.com/images/I/41y7-qWywtL._SX218_BO1,204,203,200_QL40_.jpg", Available: true}},
			{"059652692X.used_only.html",
				scraper.AmazonCoUkParsedData{Title: "Asterisk Cookbook", Price: "442.62", Image: "https://images-eu.ssl-images-amazon.com/images/I/51fXTdqAkaL._SX218_BO1,204,203,200_QL40_.jpg", Available: false}},
			{"B000Q646NA.not_a_book.html",
				scraper.AmazonCoUkParsedData{Title: "MAM Start SooTher Suitable 0-2 Months with Sterilisable Travel Case - Pack of 2, Pink", Price: "4.19", Image: "https://images-eu.ssl-images-amazon.com/images/I/41bs5q7xa1L._SY300_QL70_.jpg", Available: true}},
		}
		for _, c := range cases {
			fn, res := c.filename, c.result
			It("parsing "+fn, func() {
				fullName := filepath.Join(".", "test_data", fn)
				file, err := os.Open(fullName)
				Expect(err).To(Succeed())
				result := scraper.Parse(file)
				Expect(result).To(Equal(res))
			})
		}
	})
})

func TestScraper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "scraper")
}
