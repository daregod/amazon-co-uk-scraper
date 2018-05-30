package scraper_test

import (
	"os"
	"path/filepath"

	"github.com/daregod/amazon-co-uk-scraper/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Parser", func() {
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
	Describe("Parse", func() {
		// cases is test-wide data, stored in suite file
		for _, c := range cases {
			fn, res := c.fileName, c.result
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
