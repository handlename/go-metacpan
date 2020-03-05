package metacpan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributionName(t *testing.T) {
	d := Distribution{ID: "HTML-Restrict"}
	assert.Equal(t, "HTML::Restrict", d.Name())
}

func TestDistributionURL(t *testing.T) {
	d := Distribution{ID: "HTML-Restrict"}
	assert.Equal(t, "https://metacpan.org/pod/HTML::Restrict", d.URL())
}

func TestSearchDistribution(t *testing.T) {
	found, err := SearchDistribution("HTML-Restrict")
	assert.NoError(t, err)
	assert.Greater(t, len(found), 0)
}
