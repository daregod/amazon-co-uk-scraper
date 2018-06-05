package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	userAgent string
}

var (
	cfg config
)

func init() {
	flag.StringVar(&cfg.userAgent, "user-agent", "Ad-Blocker/1 CFNetwork/758.5.3 Darwin/15.6.0", "set request UserAgent") //nolint: lll
}

func main() {
	flag.Parse()
	urls := flag.Args()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("output to:", dir)

	for _, url := range urls {
		fileName, err := checkURL(url)
		if err != nil {
			fmt.Printf("ERROR CHECKING %s (%s)", url, err)
			continue
		}
		fmt.Println("processing", fileName)
		body, err := fetch(url)
		if err != nil {
			fmt.Printf("ERROR FETCHING %s (%s)", url, err)
			continue
		}
		err = ioutil.WriteFile(fileName, body, 0x888)
		if err != nil {
			fmt.Printf("ERROR WRITING %s (%s)", fileName, err)
			continue
		}
	}
}

func fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", cfg.userAgent)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	out := make([]byte, 0, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := res.Body.Read(buf)
		if err == nil || err != io.EOF {
			out = append(out, buf[:n]...)
			continue
		}
		break
	}

	return out, nil
}

func checkURL(url string) (string, error) {
	if !strings.Contains(url, "amazon.co.uk/gp/product/") {
		return "", fmt.Errorf("%s is not in allowed. fetch only amazon.co.uk/gp/product/<productID>", url) //nolint: lll
	}
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1] + ".html"
	return fileName, nil
}
