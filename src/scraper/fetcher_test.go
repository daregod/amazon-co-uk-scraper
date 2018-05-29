package scraper_test

import (
	"fmt"
	"io"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Fetcher", func() {
	It("Fetch", func() {
		r, err := scraper.Fetch("https://www.amazon.co.uk/gp/product/1509836071")
		Expect(err).To(Succeed())
		out := make([]byte, 0, 1024)
		buf := make([]byte, 1024)
		for {
			n, err := r.Read(buf)
			if err == nil || err != io.EOF {
				fmt.Println(n)
				out = append(out, buf[:n]...)
				continue
			}
			break
		}
		Expect(len(out)).ToNot(Equal(0))
	})
})
