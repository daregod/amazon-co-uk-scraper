package scraper

import (
	"fmt"
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
	return result
}

func (pd AmazonCoUkParsedData) Check() error {
	errIn := make([]string, 0, 3)
	if pd.Title == "" {
		errIn = append(errIn, "Title")
	}
	if pd.Available && pd.Price == "" {
		errIn = append(errIn, "Price")
	}
	if pd.Image == "" {
		errIn = append(errIn, "Image")
	}
	if len(errIn) > 0 {
		return fmt.Errorf("Parse trouble fields: %s", strings.Join(errIn, ", "))
	}
	return nil
}

func getAvailable(doc *goquery.Document) (available bool) {
	cleanup := func(text string) bool {
		return strings.Contains(text, "n stock") // In stock./Only # left in stock.
	}
	filterOut := func(bl *goquery.Selection) bool {
		if bl.Length() > 0 {
			return cleanup(bl.Text())
		}
		return false
	}

	if available = filterOut(doc.Find("div#availability")); available {
		return
	}
	doc.Find("font").Each(func(i int, selection *goquery.Selection) {
		elem := selection.Get(0)
		for _, s := range elem.Attr {
			if s.Key == "color" && s.Val == "#009900" {
				available = cleanup(elem.FirstChild.Data)
			}
		}
	})
	return
}

func getPrice(doc *goquery.Document) (price string) {
	cleanup := func(text string) string {
		return strings.TrimFunc(text, func(c rune) bool {
			if strings.ContainsAny("0123456789.", string(c)) {
				return false
			}
			return true
		})
	}
	filterOut := func(bl *goquery.Selection) string {
		pr := bl.Find("span.a-color-price")
		if pr.Length() > 0 {
			return cleanup(pr.Text())
		}
		if bl.HasClass("a-color-price") {
			return cleanup(bl.Text())
		}
		return ""
	}

	if price = filterOut(doc.Find("div#buybox")); price != "" {
		return
	}
	if price = filterOut(doc.Find("div#price")); price != "" {
		return
	}
	doc.Find("b").Each(func(i int, selection *goquery.Selection) {
		elem := selection.Get(0)
		if sel := elem.FirstChild.Data; sel == "Price:" {
			price = cleanup(elem.NextSibling.Data)
		}
	})
	if price != "" {
		return
	}
	if price = filterOut(doc.Find("span#priceblock_ourprice")); price != "" {
		return
	}
	return
}

func getImage(doc *goquery.Document) (image string) {
	filterOut := func(bl *goquery.Selection, attr string) string {
		if bl.Length() > 0 {
			if imgSrc, ok := bl.Attr(attr); ok {
				return imgSrc
			}
		}
		return ""
	}

	if image = filterOut(doc.Find("img#imgBlkFront"), "src"); image != "" {
		return
	}
	if image = filterOut(doc.Find("img#landingImage"), "src"); image != "" {
		return
	}
	if image = filterOut(doc.Find("img#detailImg"), "src"); image != "" {
		return
	}
	if image = filterOut(doc.Find("img#main-image"), "data-midres-replacement"); image != "" {
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
	if title = filterOut(doc.Find("b#product-title")); title != "" {
		return
	}
	if title = filterOut(doc.Find("h1#title")); title != "" {
		return
	}
	return
}
