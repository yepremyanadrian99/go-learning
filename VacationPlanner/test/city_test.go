package test

import (
	"github.com/golang/mock/gomock"
	"strconv"
	"testing"
	"vacation_planner/mocks"
	"vacation_planner/models"
)

func TestNewCity(t *testing.T) {
	expectedName := "Test City"
	expectedTempC := []float64{1, 2, 2.5, 3}
	city := models.NewCity(1, expectedName, expectedTempC, true, false)

	t.Run("Testing name", func(t *testing.T) {
		got := city.Name()
		if got != expectedName {
			t.Errorf("got %s, want %s", got, expectedName)
		}
	})
	t.Run("Testing TempC", func(t *testing.T) {
		got := city.TempC()
		want := expectedTempC
		for i, temp := range got {
			if temp != want[i] {
				t.Errorf("got %f, want %f", temp, want[i])
			}
		}
	})
}

func TestFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := mocks.NewMockDataReader(ctrl)

	tests := []struct {
		name      string
		responses []models.CityTemp
		query     models.CityQuery
		want      []models.CityTemp
	}{
		{
			name:      "list all",
			responses: []models.CityTemp{brCity(t, 1), srCity(t, 2)},
			query:     createQuery(t, true, false, false, 1, ""),
			want:      []models.CityTemp{brCity(t, 1), srCity(t, 2)},
		},
		{
			name:      "beach ready only",
			responses: []models.CityTemp{brCity(t, 1), srCity(t, 2)},
			query:     createQuery(t, false, true, false, 1, ""),
			want:      []models.CityTemp{brCity(t, 1)},
		},
		{
			name:      "beach ready only with name",
			responses: []models.CityTemp{brCityWithName(t, 1, "name"), brCityWithName(t, 2, "flan"), srCityWithName(t, 3, "name")},
			query:     createQuery(t, false, true, false, 1, "name"),
			want:      []models.CityTemp{brCityWithName(t, 1, "name")},
		},
		{
			name:      "ski ready only",
			responses: []models.CityTemp{brCity(t, 1), srCity(t, 2)},
			query:     createQuery(t, false, false, true, 1, ""),
			want:      []models.CityTemp{srCity(t, 2)},
		},
		{
			name:      "ski ready only with name",
			responses: []models.CityTemp{srCityWithName(t, 1, "name"), srCityWithName(t, 2, "flan"), brCityWithName(t, 3, "name")},
			query:     createQuery(t, false, false, true, 1, "name"),
			want:      []models.CityTemp{srCityWithName(t, 1, "name")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockReader.EXPECT().ReadData().Return(test.responses, nil).Times(1)
			cities, err := mockReader.ReadData()
			if err != nil {
				t.Fatal("Error reading cities: ", err)
			}

			result := models.FilterCities(cities, test.query)

			if len(result) != len(test.want) {
				t.Fatalf("\nExpected: %d results\nGot: %d results", len(test.want), len(result))
			}

			month := test.query.Month()

			for i, w := range test.want {
				if w.Id() != result[i].Id() ||
					w.Name() != result[i].Name() ||
					w.BeachVacationReady(month) != result[i].BeachVacationReady(month) ||
					w.SkiVacationReady(month) != result[i].SkiVacationReady(month) {
					t.Errorf("\nExpected: %d\nGot: %d", w, result[i])
				}
			}
		})
	}
}

func TestSorting(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := mocks.NewMockDataReader(ctrl)

	tests := []struct {
		name      string
		responses []models.CityTemp
		sort      models.CitySort
		want      []models.CityTemp
	}{
		{
			name:      "sort by id asc",
			responses: []models.CityTemp{srCity(t, 2), brCity(t, 1), brCity(t, 3)},
			sort:      createSort(t, "id", "asc"),
			want:      []models.CityTemp{brCity(t, 1), srCity(t, 2), brCity(t, 3)},
		},
		{
			name:      "sort by id desc",
			responses: []models.CityTemp{srCity(t, 2), brCity(t, 1), brCity(t, 3)},
			sort:      createSort(t, "id", "desc"),
			want:      []models.CityTemp{brCity(t, 3), srCity(t, 2), brCity(t, 1)},
		},
		{
			name:      "sort by name asc",
			responses: []models.CityTemp{srCityWithName(t, 1, "b"), brCityWithName(t, 2, "c"), brCityWithName(t, 3, "a")},
			sort:      createSort(t, "name", "asc"),
			want:      []models.CityTemp{brCityWithName(t, 3, "a"), srCityWithName(t, 1, "b"), brCityWithName(t, 2, "c")},
		},
		{
			name:      "sort by name asc",
			responses: []models.CityTemp{srCityWithName(t, 1, "b"), brCityWithName(t, 2, "c"), brCityWithName(t, 3, "a")},
			sort:      createSort(t, "name", "desc"),
			want:      []models.CityTemp{brCityWithName(t, 2, "c"), srCityWithName(t, 1, "b"), brCityWithName(t, 3, "a")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockReader.EXPECT().ReadData().Return(test.responses, nil).Times(1)
			cities, err := mockReader.ReadData()
			if err != nil {
				t.Fatal("Error reading cities: ", err)
			}

			models.SortCities(cities, test.sort)

			// Doesn't matter
			month := 10

			for i, w := range test.want {
				if w.Id() != cities[i].Id() ||
					w.Name() != cities[i].Name() ||
					w.BeachVacationReady(month) != cities[i].BeachVacationReady(month) ||
					w.SkiVacationReady(month) != cities[i].SkiVacationReady(month) {
					t.Errorf("\nExpected: %d\nGot: %d", w, cities[i])
				}
			}
		})
	}
}

func createQuery(t *testing.T, all bool, beachReady bool, skiReady bool, month int, name string) models.CityQuery {
	t.Helper()
	query, err := models.NewQuery(all, beachReady, skiReady, month, name)
	if err != nil {
		t.Fatal(err.Error())
	}
	return query
}

func createSort(t *testing.T, field string, order string) models.CitySort {
	t.Helper()
	sort, err := models.NewSort(field, order)
	if err != nil {
		t.Fatal(err.Error())
	}
	return sort
}

func brCity(t *testing.T, id int) models.CityTemp {
	return brCityWithName(t, id, strconv.FormatInt(int64(id), 10))
}

func srCity(t *testing.T, id int) models.CityTemp {
	return srCityWithName(t, id, strconv.FormatInt(int64(id), 10))
}

func brCityWithName(t *testing.T, id int, name string) models.CityTemp {
	t.Helper()
	return models.NewCity(id, name, []float64{30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}, true, false)
}

func srCityWithName(t *testing.T, id int, name string) models.CityTemp {
	t.Helper()
	return models.NewCity(id, name, []float64{-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10}, false, true)
}
