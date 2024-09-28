package test

import (
	"errors"
	"reflect"
	"testing"
	"vacation_planner/models"
)

func TestNewSort(t *testing.T) {
	// Arrange
	tests := []struct {
		testName string
		field    string
		order    string
		want     models.CitySort
		wantErr  error
	}{
		{
			testName: "create id asc sort",
			field:    "id",
			order:    "asc",
			want:     newSort(t, "id", "asc"),
		},
		{
			testName: "create id desc sort",
			field:    "id",
			order:    "desc",
			want:     newSort(t, "id", "desc"),
		},
		{
			testName: "create name asc sort",
			field:    "name",
			order:    "asc",
			want:     newSort(t, "name", "asc"),
		},
		{
			testName: "create name desc sort",
			field:    "name",
			order:    "desc",
			want:     newSort(t, "name", "desc"),
		},
		{
			testName: "create sort with unknown field",
			field:    "unknown",
			order:    "asc",
			wantErr:  errors.New("sort must be one of: id, name"),
		},
		{
			testName: "create sort with unknown order",
			field:    "id",
			order:    "unknown",
			wantErr:  errors.New("order must be one of: asc, desc"),
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			// Act
			got, err := models.NewSort(test.field, test.order)

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

func newSort(t *testing.T, field string, order string) models.CitySort {
	t.Helper()
	sort, err := models.NewSort(field, order)
	if err != nil {
		t.Fatal(err.Error())
	}
	return sort
}
