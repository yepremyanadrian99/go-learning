package main

import (
	"city_temp/models"
	"city_temp/printers"
)

func main1() {
	cities := initCities()

	p := printers.New()
	defer p.Flush()

	p.CityHeader()

	for _, city := range *cities {
		p.CityDetails(city)
	}

}

func initCities() *[]models.CityTemp {
	return &[]models.CityTemp{
		models.NewCity(1, "London", 23, false, false),
		models.NewCity(2, "Barcelona", 30, true, false),
		models.NewCity(3, "New York", 28, true, false),
		models.NewCity(4, "St. Anton", -3, false, true),
		models.NewCity(5, "Aspen", -5, false, true),
	}
}
