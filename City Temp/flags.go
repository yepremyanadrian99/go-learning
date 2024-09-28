package main

import (
	"city_temp/models"
	"city_temp/printers"
	"city_temp/reader"
	"flag"
	"fmt"
)

func transform[T any, R any](slice []T, mapper func(t T) R) []R {
	var result []R
	for _, t := range slice {
		result = append(result, mapper(t))
	}
	return result
}

func filterBeachSki(cities []models.CityTemp, all *bool, beachReady *bool, skiReady *bool) []models.CityTemp {
	if *all {
		return cities
	}

	var result []models.CityTemp
	for _, city := range cities {
		if *beachReady && city.BeachVacationReady() || *skiReady && city.SkiVacationReady() {
			result = append(result, city)
		}
	}
	return result
}

func main() {
	all := flag.Bool("a", false, "Display all destinations")
	beachReady := flag.Bool("b", false, "Display beach ready destinations")
	skiReady := flag.Bool("s", false, "Display ski ready destinations")
	flag.Parse()

	dataReader := reader.NewDataReader()
	responses, err := dataReader.ReadData()

	if err != nil {
		fmt.Println("Error while reading data: ", err)
		return
	}

	var cities = transform(responses, models.ToCityTemp)
	cities = filterBeachSki(cities, all, beachReady, skiReady)

	printer := printers.New()
	defer printer.Flush()

	printer.CityHeader()
	for _, city := range cities {
		printer.CityDetails(city)
	}
}
