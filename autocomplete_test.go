package metacpan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchAutocomplete(t *testing.T) {
	hits, err := SearchAutocomplete("HTML")
	assert.NoError(t, err)
	assert.Greater(t, len(hits), 5)
	assert.Contains(t, hits[0].URL(), "https://metacpan.org/pod/")
}
