package totus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReference_GeoPOI(t *testing.T) {
	client, err := NewTotus("", "", "")
	assert.NoError(t, err)
	ref := client.Reference()
	pois, err := ref.GeoPOI(
		NewGeoPOISearch().
			WithGeoHash("69y7pkxfc").
			WithWhat("shop").
			WithDistance(1000.0).
			WithLimit(10))
	assert.NoError(t, err)
	assert.Len(t, pois, 10)
}

func TestReference_IP(t *testing.T) {
	client, err := NewTotus("", "", "")
	assert.NoError(t, err)
	ref := client.Reference()

	resp, err := ref.NetIP4("8.8.8.8")
	assert.NoError(t, err)
	assert.Equal(t, "8.8.8.8", resp.IP4())
	assert.Equal(t, "9q9htvvm81jd", resp.GH())

	resp, err = ref.NetIP6("2001:4860:4860::8888")
	assert.NoError(t, err)
	assert.Equal(t, "2001:4860:4860::8888", resp.IP6())
}
