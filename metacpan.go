package metacpan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// API host
const host = "https://api.metacpan.org"

// APIs
const (
	APISearchDistribution = "/v0/distribution/_search"
)

// The Distribution defines type of a distribution.
type Distribution struct {
	ID string `json:"_id"`
}

// Name returns module name.
func (d Distribution) Name() string {
	return strings.Replace(d.ID, "-", "::", -1)
}

// URL returns url on metacpan.
func (d Distribution) URL() string {
	return fmt.Sprintf("https://metacpan.org/pod/%s", d.Name())
}

// The DistributionHits defines search results.
type DistributionHits struct {
	Hits []Distribution `json:"hits"`
}

// The SearchDistributionResult defines result of search distribution.
type SearchDistributionResult struct {
	Hits DistributionHits `json:"hits"`
}

// SearchDistribution search distribution by query and returns hits.
func SearchDistribution(q string) ([]Distribution, error) {
	body, err := request(fmt.Sprintf("%s%s?q=%s", host, APISearchDistribution, q))

	fmt.Println(string(body))

	if err != nil {
		fmt.Println("error at request search distribution.")
		return nil, err
	}

	var result SearchDistributionResult
	err = json.Unmarshal(body, &result)

	if err != nil {
		fmt.Println("error at unmarshal response of search distribution.")
		return nil, err
	}

	return result.Hits.Hits, nil
}

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
