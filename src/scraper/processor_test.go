package scraper_test

import (
	"encoding/json"
	"fmt"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Processor", func() {
	PIt("Process", func() {
		jd := scraper.ProcessUrls(urls)
		q, _ := json.MarshalIndent(jd, "", "\t")
		fmt.Println(string(q))

		Expect(true).To(BeTrue())
	})
})

var urls []string = []string{
	"https://www.amazon.co.uk/gp/product/B000Q646NA",
	"https://www.amazon.co.uk/gp/product/1787125645",
	"https://www.amazon.co.uk/gp/product/059652692X",
	"https://www.amazon.co.uk/gp/product/1509836071",
}
