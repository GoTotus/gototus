# GoTotus.com APIs' Golang library

## Basic Usage

`TOTUS_KEY` environment variable will be used to pick the api
key ([create one here](https://gototus.com/console/apikeys))

```
    func main() {
        t, err := totus.NewTotus("", "", "")
        ref := t.Reference()
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
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
    }

```

it will print:

```
Any shop nearby:
{
    "dist": 71.6,
    "gh": "69y7pkx5r3",
    "id": 4675113766,
    "info": {
        "addr:city": "Ciudad Autónoma de Buenos Aires",
        "addr:country": "AR",
        "addr:street": "Avenida Corrientes",
        "name": "Maxikiosko",
        "shop": "kiosk"
    },
    "lat": -34.60362,
    "lon": -58.3824
}
{
    "dist": 84,
    "gh": "69y7ps83ms",
    "id": 12179098601,
    "info": {
        "addr:housenumber": "999",
        "addr:street": "Avenida Presidente Roque Sáenz Peña",
        "name": "I Love Gifts",
        "shop": "gift"
    },
    "lat": -34.60395,
    "lon": -58.38076
}
```

## Examples

For further examples, check the `examples/` folder in this project.
Or a public copy at the [GitHub Website](https://github.com/GoTotus/gototus/tree/main/examples).

## Manuals

For detailed manuals about Totus please check: [docs.gototus.com](https://docs.gototus.com)

## Installing

`go get github.com/GoTotus/gototus`

## Building examples

`make`

for building, you will need to have a functioning Totus API key in the envvar `TOTUS_KEY`.
