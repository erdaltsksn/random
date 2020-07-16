package random

import (
	"math/rand"
	"time"
)

// NameStr is a struct hold data.
type NameStr struct {
	Name    string
	Gender  string
	Country string
}

// Name generates a random name according specified gender and country.
func Name(gender, country string) string {
	var dataWithGender []NameStr
	var data []NameStr

	if gender != "" {
		for _, v := range NameList {
			if v.Gender == gender {
				dataWithGender = append(dataWithGender, v)
			}
		}
	}

	if len(dataWithGender) == 0 {
		dataWithGender = NameList
	}

	if country != "" {
		for _, v := range dataWithGender {
			if v.Country == country {
				data = append(data, v)
			}
		}
	}

	if len(data) == 0 {
		data = NameList
	}

	rand.Seed(time.Now().Unix())

	return data[rand.Intn(len(data))].Name
}
