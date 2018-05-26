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
	avail, price := getAvailableAndPrice(doc)
	img := getImage(doc)
	title := getTitle(doc)
	result := AmazonCoUkParsedData{
		Title:     title,
		Price:     price,
		Image:     img,
		Available: avail,
	}

	//	q, _ := json.MarshalIndent(result, "", "\t")
	//	fmt.Println(string(q))
	return result
}

func getAvailableAndPrice(doc *goquery.Document) (available bool, price string) {
	buyBox := doc.Find("div#buybox")

	avail := buyBox.Find("div#availability")
	if avail.Length() > 0 {
		available = strings.Contains(avail.Text(), "In stock")
	}

	pr := buyBox.Find("span.a-color-price")
	if pr.Length() > 0 {
		price = strings.TrimFunc(pr.Text(), func(c rune) bool {
			if strings.ContainsAny("0123456789.", string(c)) {
				return false
			}
			return true
		})
	}
	return
}

func getImage(doc *goquery.Document) (image string) {
	img := doc.Find("img#imgBlkFront")
	if img.Length() > 0 {
		if imgSrc, ok := img.Attr("src"); ok {
			image = imgSrc
		}
	}
	return
}

func getTitle(doc *goquery.Document) (title string) {
	ttl := doc.Find("span#productTitle")
	if ttl.Length() > 0 {
		title = ttl.Text()
	}
	return
}
