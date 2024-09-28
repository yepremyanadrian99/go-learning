package models

import (
	"errors"
)

type CityQuery interface {
	All() bool
	Beach() bool
	Ski() bool
	Month() int
	Name() string
}

func NewQuery(all bool, beach bool, ski bool, month int, name string) (CityQuery, error) {
	q := &cityQuery{
		all:   all,
		beach: beach,
		ski:   ski,
		month: month,
		name:  name,
	}
	if err := q.validate(); err != nil {
		return nil, err
	}
	return q, nil
}

type cityQuery struct {
	all   bool
	beach bool
	ski   bool
	month int
	name  string
}

func (q cityQuery) All() bool    { return q.all }
func (q cityQuery) Beach() bool  { return q.beach }
func (q cityQuery) Ski() bool    { return q.ski }
func (q cityQuery) Month() int   { return q.month }
func (q cityQuery) Name() string { return q.name }

func (q cityQuery) validate() error {
	if q.month < 1 || q.month > 12 {
		return errors.New("month must be between 1 and 12")
	}

	return nil
}
