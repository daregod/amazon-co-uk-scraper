package parser

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AmazonCoUkParsedData struct {
	Title     string `json:"title"`
	Price     string `json:"price"`
	Image     string `json:"image"`
	Available bool   `json:"available"`
}
type AmazonCoUkBulkData struct {
	Url  string               `json:"url"`
	Meta AmazonCoUkParsedData `json:"meta"`
}

func Parse(r io.Reader) AmazonCoUkParsedData {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	result := AmazonCoUkParsedData{}
	buyBox := doc.Find("div#buybox")

	avail := buyBox.Find("div#availability")
	if avail.Length() > 0 {
		result.Available = strings.Contains(avail.Text(), "In stock")
	}

	price := buyBox.Find("span.a-color-price")
	if price.Length() > 0 {
		result.Price = strings.TrimFunc(price.Text(), func(c rune) bool {
			if strings.ContainsAny("0123456789.", string(c)) {
				return false
			}
			return true
		})
	}

	//	q, _ := json.MarshalIndent(result, "", "\t")
	//	fmt.Println(string(q))
	return result
}
