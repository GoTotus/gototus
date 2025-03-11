package totus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReference_GeoPOI(t *testing.T) {
	client, err := NewTotus("", "", "")
	assert.NoError(t, err)
	ref := client.Reference()

	gh := "69y7pkxfc"
	what := "shop"
	dist := 1000.0
	limit := 10
	params := GeoPOIParams{
		GH:       &gh,
		What:     &what,
		Distance: &dist,
		Limit:    &limit,
	}
	pois, err := ref.GeoPOI(params)
	assert.NoError(t, err)
	assert.Len(t, pois, 10)
}

func TestReference_IP(t *testing.T) {
	client, err := NewTotus("", "", "")
	assert.NoError(t, err)
	ref := client.Reference()

	resp, err := ref.IP("8.8.8.8", "")
	if err != nil {
		t.Fatalf("IP() error = %v", err)
	}
	if resp.IP4() != "8.8.8.8" || resp.GH() != "9q9htvvm81jd" {
		t.Errorf("IP() = %v, want ip4=8.8.8.8, gh=9q9htvvm81jd", resp)
	}
}
