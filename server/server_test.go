package server_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "server")
}

var _ = Describe("Server", func() {
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
})
