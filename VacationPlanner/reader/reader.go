package reader

import (
	"encoding/json"
	"fmt"
	"os"
	"vacation_planner/models"
	"vacation_planner/utils"
)

//go:generate mockgen -destination=../mocks/mock_reader.go -package=mocks vacation_planner/reader DataReader
type DataReader interface {
	ReadData() ([]models.CityTemp, error)
}

func NewDataReader() DataReader {
	return &reader{
		path: "./data/cities.json",
	}
}

type reader struct {
	path string
}

func (reader *reader) ReadData() ([]models.CityTemp, error) {
	file, err := os.ReadFile(reader.path)
	if err != nil {
		fmt.Println("Error while reading file: ", reader.path, err)
		return nil, err
	}

	var responses []models.CityResponse
	err = json.Unmarshal(file, &responses)
	if err != nil {
		fmt.Println("Error while unmarshalling file: ", reader.path, err)
		return nil, err
	}

	cityTemps := utils.Transform(responses, models.ToCityTemp)

	return cityTemps, nil
}
