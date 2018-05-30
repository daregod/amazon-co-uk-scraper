package scraper_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"

	"github.com/gocolly/colly"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Processor", func() {
	var ts *httptest.Server
	BeforeEach(func() {
		ts = newTestServer()
	})
	It("Process", func() {
		proc := scraper.AmazonCoUkProcessor{
			Collector: colly.NewCollector(),
		}
		baseUrl := ts.URL
		jd := proc.ProcessUrls(urlsFromCases(baseUrl))
		Expect(jd).To(ConsistOf(bulkFromCases(baseUrl)))
	})
})

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	// cases is test-wide data, stored in suite file
	for _, c := range cases {
		fn, hp := c.fileName, c.handlePath
		mux.HandleFunc(hp, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fullName := filepath.Join(".", "test_data", fn)
			file, err := ioutil.ReadFile(fullName)
			Expect(err).To(Succeed())
			w.Write(file)
		})
	}
	return httptest.NewServer(mux)
}
