package scraper

import (
	"fmt"
	"io"
	"net/http"
)

func Fetch(url string) (io.Reader, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Ad-Blocker/1 CFNetwork/758.5.3 Darwin/15.6.0")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}
