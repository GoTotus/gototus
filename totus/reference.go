package totus

import (
	"fmt"
	"net/url"
)

// Reference provides methods for accessing reference data from the Totus API
type Reference struct {
	totus *Totus
}

// GeoPOIParams holds parameters for the GeoPOI method
type GeoPOIParams struct {
	Lat      *float64
	Lon      *float64
	GH       *string
	What     *string
	Distance *float64
	Filter   map[string]string
	Limit    *int
}

// GeoPOI fetches points of interest based on the provided parameters
func (r *Reference) GeoPOI(params GeoPOIParams) ([]POI, error) {
	q := url.Values{}
	if params.Lat != nil {
		q.Add("lat", fmt.Sprintf("%f", *params.Lat))
	}
	if params.Lon != nil {
		q.Add("lon", fmt.Sprintf("%f", *params.Lon))
	}
	if params.GH != nil {
		q.Add("gh", *params.GH)
	}
	if params.What != nil {
		q.Add("what", *params.What)
	}
	if params.Distance != nil {
		q.Add("dist", fmt.Sprintf("%f", *params.Distance))
	}
	if params.Filter != nil {
		for k, v := range params.Filter {
			q.Add("filter", fmt.Sprintf("%s=%s", k, v))
		}
	}
	if params.Limit != nil {
		q.Add("limit", fmt.Sprintf("%d", *params.Limit))
	}

	var pois []POI
	err := r.totus.makeRequest("GET", "/ref/geo/poi", q, nil, &pois)
	if err != nil {
		return nil, err
	}
	return pois, nil
}

// IP fetches IP information for the given IPv4 or IPv6 address
func (r *Reference) IP(ip4, ip6 string) (*IPData, error) {
	q := url.Values{}
	if ip4 != "" {
		q.Add("ip4", ip4)
	}
	if ip6 != "" {
		q.Add("ip6", ip6)
	}

	var resp IPData
	err := r.totus.makeRequest("GET", "/ref/ip", q, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
