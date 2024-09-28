package models

import (
	"errors"
	"strings"
)

var allowedSorts = []string{"id", "name"}
var allowedOrders = []string{"asc", "desc"}

type CitySort interface {
	Field() string
	Order() string
}

func NewSort(field string, order string) (CitySort, error) {
	s := &citySort{
		field: field,
		order: order,
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return s, nil
}

type citySort struct {
	field string
	order string
}

func (s citySort) Field() string { return s.field }
func (s citySort) Order() string { return s.order }

func (s citySort) validate() error {
	if !contains(allowedSorts, s.field) {
		return errors.New("sort must be one of: " + strings.Join(allowedSorts, ", "))
	}

	if !contains(allowedOrders, s.order) {
		return errors.New("order must be one of: " + strings.Join(allowedOrders, ", "))
	}

	return nil
}

func contains(strings []string, value string) bool {
	for _, s := range strings {
		if value == s {
			return true
		}
	}
	return false
}
