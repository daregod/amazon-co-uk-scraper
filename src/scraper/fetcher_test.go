package scraper_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/daregod/amazon-co-uk-scraper/src/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Fetcher", func() {
	XIt("Fetch", func() {
		r, err := scraper.Fetch("https://www.amazon.co.uk/gp/product/059652692X")
		Expect(err).To(Succeed())
		var buf bytes.Buffer
		body := io.TeeReader(r, &buf)
		pd := scraper.Parse(body)
		q, _ := json.MarshalIndent(pd, "", "\t")
		fmt.Println(string(q))
		ioutil.WriteFile("./test_data/1.html", buf.Bytes(), 0x888)
		//		out := make([]byte, 0, 1024)
		//		buf := make([]byte, 1024)
		//		for {
		//			n, err := r.Read(buf)
		//			if err == nil || err != io.EOF {
		//				out = append(out, buf[:n]...)
		//				continue
		//			}
		//			break
		//		}
		//		Expect(len(out)).ToNot(Equal(0))
	})
})
