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
	pois, err := ref.GeoPOI(
		totus.NewGeoPOISearch().
			WithGeoHash("69y7pkxfc").
			WithWhat("shop").
			WithDistance(1000.0).
			WithLimit(2))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

	fmt.Println("Any shop nearby, but providing lat/lon instead of geohash:")
	pois, err = ref.GeoPOI(
		totus.NewGeoPOISearch().
			WithLat(-34.60362).
			WithLon(-58.3824).
			WithWhat("shop").
			WithDistance(1000.0).
			WithLimit(2))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

	fmt.Println("Only bookshops, 2km around:")
	pois, err = ref.GeoPOI(
		totus.NewGeoPOISearch().
			WithLat(-34.60362).
			WithLon(-58.3824).
			WithWhat("shop").
			WithDistance(2000.0).
			WithLimit(2))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

	fmt.Println("Only bookshops, 2km around, name includes the word 'libro' in any case:")
	pois, err = ref.GeoPOI(
		totus.NewGeoPOISearch().
			WithLat(-34.60362).
			WithLon(-58.3824).
			WithWhat("shop").
			WithDistance(2000.0).
			AddFilter("shop", "books").
			AddFilter("name", "~*libro*").
			WithLimit(2))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, p := range pois {
		fmt.Println(p)
	}

}
