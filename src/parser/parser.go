package parser

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func Parse(r io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	return doc.Find("span.a-color-price").First().Text()
}
