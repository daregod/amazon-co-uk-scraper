package scraper

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
	avail := getAvailable(doc)
	price := getPrice(doc)
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

func getAvailable(doc *goquery.Document) (available bool) {
	avail := doc.Find("div#availability")
	if avail.Length() > 0 {
		available = strings.Contains(avail.Text(), "n stock") // In stock./Only # left in stock.
	}

	return
}

func getPrice(doc *goquery.Document) (price string) {
	filterOut := func(bl *goquery.Selection) string {
		pr := bl.Find("span.a-color-price")
		if pr.Length() > 0 {
			return strings.TrimFunc(pr.Text(), func(c rune) bool {
				if strings.ContainsAny("0123456789.", string(c)) {
					return false
				}
				return true
			})
		}
		return ""
	}

	if price = filterOut(doc.Find("div#buybox")); price != "" {
		return
	}
	if price = filterOut(doc.Find("div#price")); price != "" {
		return
	}
	return
}

func getImage(doc *goquery.Document) (image string) {
	filterOut := func(bl *goquery.Selection) string {
		if bl.Length() > 0 {
			if imgSrc, ok := bl.Attr("src"); ok {
				return imgSrc
			}
		}
		return ""
	}
	if image = filterOut(doc.Find("img#imgBlkFront")); image != "" {
		return
	}
	if image = filterOut(doc.Find("img#landingImage")); image != "" {
		return
	}
	return
}

func getTitle(doc *goquery.Document) (title string) {
	filterOut := func(bl *goquery.Selection) string {
		if bl.Length() > 0 {
			return strings.Trim(bl.Text(), " \r\n\t")
		}
		return ""
	}
	if title = filterOut(doc.Find("span#productTitle")); title != "" {
		return
	}
	return
}
