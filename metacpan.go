package metacpan

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// API host
const host = "https://api.metacpan.org"

// APIs
const (
	APISearchDistribution = "/v0/distribution/_search"
)

func request(urlStr string) ([]byte, error) {
	resp, err := http.DefaultClient.Get(urlStr)

	if err != nil {
		fmt.Println("error at request", urlStr)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error at read response body.")
		return nil, err
	}

	return body, nil
}
