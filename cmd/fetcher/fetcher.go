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

type Config struct {
	UA string
}

var (
	cfg Config
)

func init() {
	flag.StringVar(&cfg.UA, "user-agent", "Ad-Blocker/1 CFNetwork/758.5.3 Darwin/15.6.0", "set request UserAgent")
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
		fileName, err := checkUrl(url)
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
		ioutil.WriteFile(fileName, body, 0x888)
	}
}

func fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", cfg.UA)

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

func checkUrl(url string) (string, error) {
	if !strings.Contains(url, "amazon.co.uk/gp/product/") {
		return "", fmt.Errorf("%s is not in allowed. fetch only amazon.co.uk/gp/product/<productID>", url)
	}
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1] + ".html"
	return fileName, nil
}
