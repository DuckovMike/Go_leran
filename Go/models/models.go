package models

import (
	"time"
)

type User struct {
	Id         int
	Name       string
	Prof       string
	Rate       float64
	Year       int
	Month      int
	Appearance Appearence
}

type Appearence struct {
	Dates map[int]int
}

type DayType int

const (
	Workday DayType = iota
	Weekend
	ShortDay
)

func (u *User) CreateDefaultAppearence(year int, month int, part int) {
	dates := map[int]int{}
	if part == 1 {
		for i := 1; i <= 15; i++ {
			iWorkDay := int(time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.UTC).Weekday())
			switch iWorkDay {
			case 0, 6:
				dates[i] = int(Weekend)
			default:
				dates[i] = int(Workday)
			}

		}
	} else {
		lastDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
		for i := 16; i <= lastDay.Day(); i++ {
			iWorkDay := int(time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.UTC).Weekday())
			switch iWorkDay {
			case 0, 6:
				dates[i] = 1
			default:
				dates[i] = 0
			}
		}
	}
	u.Appearance = Appearence{Dates: dates}
	u.Month = month
	u.Year = year
}

func CreateUser(id int, name string, prof string, rate float64) *User {
	return &User{
		Id:         id,
		Name:       name,
		Prof:       prof,
		Rate:       rate,
		Appearance: Appearence{},
	}
}
