package models

import (
	"sort"
	"strings"
	"vacation_planner/utils"
)

const (
	beachThreshold float64 = 22
	skiThreshold   float64 = -2
)

type CityTemp interface {
	Id() int
	Name() string
	TempC() []float64
	TempF() []float64
	BeachVacationReady(int) bool
	SkiVacationReady(int) bool
}

type CityResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	TempC       []float64 `json:"temp_c"`
	HasBeach    bool      `json:"has_beach"`
	HasMountain bool      `json:"has_mountain"`
}

type city struct {
	id          int
	name        string
	tempC       []float64
	hasBeach    bool
	hasMountain bool
}

func NewCity(id int, name string, tempC []float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{id: id, name: name, tempC: tempC, hasBeach: hasBeach, hasMountain: hasMountain}
}

func (city city) Id() int {
	return city.id
}

func (city city) Name() string {
	return city.name
}

func (city city) TempC() []float64 {
	return city.tempC
}

func (city city) TempF() []float64 {
	return utils.Transform(city.tempC, func(tempC float64) float64 {
		return (tempC * 9 / 5) + 32
	})
}

func (city city) BeachVacationReady(month int) bool {
	return city.hasBeach && city.tempC[month-1] >= beachThreshold
}

func (city city) SkiVacationReady(month int) bool {
	return city.hasMountain && city.tempC[month-1] <= skiThreshold
}

func ToCityTemp(response CityResponse) CityTemp {
	return response.ToCityTemp()
}

func (response CityResponse) ToCityTemp() CityTemp {
	return city{
		id:          response.Id,
		name:        response.Name,
		tempC:       response.TempC,
		hasBeach:    response.HasBeach,
		hasMountain: response.HasMountain,
	}
}

func FilterCities(cities []CityTemp, q CityQuery) []CityTemp {
	var result []CityTemp

	skiBeachMatch := func(city CityTemp) bool {
		return q.All() || q.Ski() && city.SkiVacationReady(q.Month()) || q.Beach() && city.BeachVacationReady(q.Month())
	}
	nameMatch := func(city CityTemp) bool {
		return q.Name() == "" || strings.Contains(strings.ToLower(city.Name()), strings.ToLower(q.Name()))
	}

	for _, city := range cities {
		if skiBeachMatch(city) && nameMatch(city) {
			result = append(result, city)
		}
	}
	return result
}

func SortCities(cities []CityTemp, s CitySort) {
	sort.Slice(cities, func(i, j int) bool {
		if s.Field() == "id" {
			return orderById(cities[i].Id(), cities[j].Id(), s.Order())
		} else if s.Field() == "name" {
			return orderByName(cities[i].Name(), cities[j].Name(), s.Order())
		}
		return false
	})
}

func orderById(v1, v2 int, order string) bool {
	if order == "asc" {
		return v1 <= v2
	} else {
		return v1 >= v2
	}
}

func orderByName(v1, v2 string, order string) bool {
	v1 = strings.ToLower(v1)
	v2 = strings.ToLower(v2)
	if order == "asc" {
		return v1 <= v2
	} else {
		return v1 >= v2
	}
}
