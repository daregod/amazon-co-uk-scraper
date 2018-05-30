package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScrapeServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cmd/scrape-server")
}

var _ = Describe("Scrape-server", func() {
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
})
