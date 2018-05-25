package parser_test

import (
	"os"
	"testing"

	"github.com/daregod/amazon-co-uk-scraper/src/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Parser", func() {
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
	It("Parse", func() {
		div, err := os.Open("./test_data/div_new_in_stock.txt")
		Expect(err).To(Succeed())
		result := parser.Parse(div)
		Expect(result.Price).To(Equal(`8.49`))
		Expect(result.Available).To(BeTrue())

		div, err = os.Open("./test_data/div_new_and_used.txt")
		Expect(err).To(Succeed())
		result = parser.Parse(div)
		Expect(result.Price).To(Equal(`41.99`))
		Expect(result.Available).To(BeTrue())

		div, err = os.Open("./test_data/div_used_only.txt")
		Expect(err).To(Succeed())
		result = parser.Parse(div)
		Expect(result.Price).To(Equal(`442.62`))
		Expect(result.Available).To(BeFalse())
	})
})

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "parser")
}

var testData []byte = []byte(``)
