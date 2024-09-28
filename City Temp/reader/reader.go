package reader

import (
	"city_temp/models"
	"encoding/json"
	"fmt"
	"os"
)

type DataReader interface {
	ReadData() ([]models.CityResponse, error)
}

func NewDataReader() DataReader {
	return &reader{
		path: "./data/cities.json",
	}
}

type reader struct {
	path string
}

func (reader *reader) ReadData() ([]models.CityResponse, error) {
	file, err := os.ReadFile(reader.path)
	if err != nil {
		fmt.Println("Error while reading file: ", reader.path, err)
		return nil, err
	}

	var data []models.CityResponse
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error while unmarshalling file: ", reader.path, err)
		return nil, err
	}

	return data, nil
}
