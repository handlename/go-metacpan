package metacpan

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// hosts
const (
	apiHost  = "fastapi.metacpan.org"
	htmlHost = "metacpan.org"
)

// APIs
const (
	APISearchAutocomplete = "/v1/search/autocomplete"
	APISearchDistribution = "/v1/distribution/_search"
)

func request(urlStr string) ([]byte, error) {
	resp, err := http.DefaultClient.Get(fmt.Sprintf("https://%s%s", apiHost, urlStr))

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
