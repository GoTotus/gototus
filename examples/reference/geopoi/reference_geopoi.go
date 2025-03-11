package main

import (
	"fmt"
	"github.com/GoTotus/gototus/totus"
)

func main() {
	t, err := totus.NewTotus("", "", "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	ref := t.Reference()

	fmt.Println("Any shop nearby:")
	gh := "69y7pkxfc"
	what := "shop"
	dist := 1000.0
	limit := 2
	pois, err := ref.GeoPOI(totus.GeoPOIParams{
		GH:       &gh,
		What:     &what,
		Distance: &dist,
		Limit:    &limit,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

	fmt.Println("Any shop nearby, but providing lat/lon instead of geohash:")
	lat := -34.60362
	lon := -58.3824
	pois, err = ref.GeoPOI(totus.GeoPOIParams{
		Lat:   &lat,
		Lon:   &lon,
		What:  &what,
		Limit: &limit,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

	fmt.Println("Only bookshops, 2km around:")
	dist = 2000.0
	pois, err = ref.GeoPOI(totus.GeoPOIParams{
		GH:       &gh,
		What:     &what,
		Distance: &dist,
		Filter:   map[string]string{"shop": "books"},
		Limit:    &limit,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

	fmt.Println("Only bookshops, 2km around, name includes the word 'libro' in any case:")
	dist = 2000.0
	what = "shop"
	pois, err = ref.GeoPOI(totus.GeoPOIParams{
		GH:       &gh,
		What:     &what,
		Distance: &dist,
		Filter:   map[string]string{"shop": "books", "name": "~*libro*"},
		Limit:    &limit,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

}
