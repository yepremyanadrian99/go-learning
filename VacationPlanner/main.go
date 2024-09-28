package main

import (
	"flag"
	"log"
	"vacation_planner/models"
	"vacation_planner/printer"
	"vacation_planner/reader"
)

func main() {
	query, sort, err := readQueryAndSort()
	if err != nil {
		log.Fatal(err)
	}

	dataReader := reader.NewDataReader()

	cities, err := dataReader.ReadData()
	if err != nil {
		log.Fatal(err)
	}

	cities = models.FilterCities(cities, query)
	models.SortCities(cities, sort)

	printer := printers.New()
	defer printer.Flush()

	printer.CityHeader()
	for _, city := range cities {
		printer.CityDetails(city, query.Month())
	}
}

func readQueryAndSort() (models.CityQuery, models.CitySort, error) {
	allFlag := flag.Bool("all", false, "To include all destinations")
	skiFlag := flag.Bool("ski", false, "To include ski ready destinations")
	beachFlag := flag.Bool("beach", false, "To include beach ready destinations")
	monthFlag := flag.Int("month", 0, "Specify month of vacation")
	nameFlag := flag.String("name", "", "Search for a destination name")
	sortFlag := flag.String("sort", "id", "Specify sort field (id | name)")
	orderFlag := flag.String("order", "asc", "Specify sort order (asc | desc)")

	flag.Parse()

	q, err := models.NewQuery(*allFlag, *skiFlag, *beachFlag, *monthFlag, *nameFlag)
	if err != nil {
		return nil, nil, err
	}

	s, err := models.NewSort(*sortFlag, *orderFlag)
	if err != nil {
		return nil, nil, err
	}

	return q, s, nil
}
