package test

import (
	"errors"
	"reflect"
	"testing"
	"vacation_planner/models"
)

func TestNewQuery(t *testing.T) {
	// Arrange
	tests := []struct {
		testName string
		all      bool
		beach    bool
		ski      bool
		month    int
		name     string
		want     models.CityQuery
		wantErr  error
	}{
		{
			testName: "create query successfully",
			all:      false,
			beach:    true,
			ski:      true,
			month:    1,
			name:     "name",
			want:     newQuery(t, false, true, true, 1, "name"),
		},
		{
			testName: "create query with invalid month 1",
			all:      false,
			beach:    true,
			ski:      true,
			month:    0,
			name:     "name",
			wantErr:  errors.New("month must be between 1 and 12"),
		},
		{
			testName: "create query with invalid month 2",
			all:      false,
			beach:    true,
			ski:      true,
			month:    13,
			name:     "name",
			wantErr:  errors.New("month must be between 1 and 12"),
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			// Act
			got, err := models.NewQuery(test.all, test.beach, test.ski, test.month, test.name)

			// Assert
			if test.wantErr != nil {
				if err == nil || err.Error() != test.wantErr.Error() {
					t.Errorf("\nGot error: %v\nWant error: %v\n", err, test.wantErr)
				}
				return
			}

			// Unwanted / unexpected error
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("\nGot: %v\nWant: %v\n", got, test.want)
			}
		})
	}
}

func newQuery(t *testing.T, all bool, beachReady bool, skiReady bool, month int, name string) models.CityQuery {
	t.Helper()
	query, err := models.NewQuery(all, beachReady, skiReady, month, name)
	if err != nil {
		t.Fatal(err.Error())
	}
	return query
}
