package totus

import (
	"fmt"
	"net/url"
)

// Reference provides methods for accessing reference data from the Totus API
type Reference struct {
	totus *Totus
}

type GeoPOISearch struct {
	Lat      *float64
	Lon      *float64
	GH       *string
	What     *string
	Distance *float64
	Filter   map[string]string
	Limit    *int
}

func NewGeoPOISearch() GeoPOISearch {
	return GeoPOISearch{}
}

func (g GeoPOISearch) WithLat(lat float64) GeoPOISearch {
	g.Lat = &lat
	return g
}

func (g GeoPOISearch) WithLon(lat float64) GeoPOISearch {
	g.Lon = &lat
	return g
}

func (g GeoPOISearch) WithGeoHash(gh string) GeoPOISearch {
	g.GH = &gh
	return g
}

func (g GeoPOISearch) WithWhat(what string) GeoPOISearch {
	g.What = &what
	return g
}

func (g GeoPOISearch) WithDistance(distance float64) GeoPOISearch {
	g.Distance = &distance
	return g
}

func (g GeoPOISearch) WithLimit(limit int) GeoPOISearch {
	g.Limit = &limit
	return g
}

func (g GeoPOISearch) AddFilter(key, value string) GeoPOISearch {
	if g.Filter == nil {
		g.Filter = make(map[string]string)
	}
	g.Filter[key] = value
	return g
}

func (r *Reference) GeoPOI(params GeoPOISearch) ([]POI, error) {
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

func (r *Reference) IP() (*IPData, error) {
	var resp IPData
	err := r.totus.makeRequest("GET", "/ref/ip", nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// IP fetches IP information for the given IPv4 or IPv6 address
func (r *Reference) IP4(ip4 string) (*IPData, error) {
	q := url.Values{}
	q.Add("ip4", ip4)

	var resp IPData
	err := r.totus.makeRequest("GET", "/ref/ip", q, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// IP fetches IP information for the given IPv4 or IPv6 address
func (r *Reference) IP6(ip6 string) (*IPData, error) {
	q := url.Values{}
	q.Add("ip6", ip6)

	var resp IPData
	err := r.totus.makeRequest("GET", "/ref/ip", q, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
