package metacpan

import (
	"encoding/json"
	"fmt"
	"strings"
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
	return fmt.Sprintf("https://%s/pod/%s", htmlHost, d.Name())
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
	u := fmt.Sprintf("%s?q=%s", APISearchDistribution, q)
	body, err := request(u)

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
