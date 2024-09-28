package main

import (
	"city_temp/printers"
	"city_temp/reader"
	"fmt"
	"sort"
)

func main4() {
	dataReader := reader.NewDataReader()
	cities, err := dataReader.ReadData()

	sort.Slice(cities, func(i, j int) bool {
		return cities[i].Name < cities[j].Name
	})

	if err != nil {
		fmt.Println("Error while reading data: ", err)
		return
	}

	printer := printers.New()
	defer printer.Flush()

	printer.CityHeader()
	for _, city := range cities {
		printer.CityDetails(city.ToCityTemp())
	}
}
