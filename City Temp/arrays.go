package main

import (
	"fmt"
)

func addCity(city string, cities *[2]string) {
	cities[1] = city
}

func main2() {
	cities := [2]string{"London"}

	citiesCopy := cities

	cities[1] = "Adrian"
	addCity("Bro", &citiesCopy)

	fmt.Println("Cities: ", cities)
	fmt.Println("Copy: ", citiesCopy)
}
