package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFetcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cmd/fetcher")
}

var _ = Describe("Server", func() {
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
})
