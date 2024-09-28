package models

const (
	beachThreshold float64 = 22
	skiThreshold   float64 = -2
)

type CityTemp interface {
	Id() int
	Name() string
	TempC() float64
	TempF() float64
	BeachVacationReady() bool
	SkiVacationReady() bool
}

type CityResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	TempC       float64 `json:"temp_c"`
	HasBeach    bool    `json:"has_beach"`
	HasMountain bool    `json:"has_mountain"`
}

type city struct {
	id          int
	name        string
	tempC       float64
	hasBeach    bool
	hasMountain bool
}

func NewCity(id int, name string, tempC float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{id: id, name: name, tempC: tempC, hasBeach: hasBeach, hasMountain: hasMountain}
}

func (city city) Id() int {
	return city.id
}

func (city city) Name() string {
	return city.name
}

func (city city) TempC() float64 {
	return city.tempC
}

func (city city) TempF() float64 {
	return (city.tempC * 9 / 5) + 32
}

func (city city) BeachVacationReady() bool {
	return city.hasBeach && city.tempC >= beachThreshold
}

func (city city) SkiVacationReady() bool {
	return city.hasMountain && city.tempC <= skiThreshold
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
