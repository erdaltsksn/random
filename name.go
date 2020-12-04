package random

import (
	"fmt"
)

// Person is a struct hold data.
type Person struct {
	Name    string
	Gender  string
	Country string
}

// Firstname generates a random first name according specified gender and country.
func Firstname(gender, country string) string {
	randSeed()

	var dataWithGender []Person
	var data []Person

	if gender != "" {
		for _, v := range FirstNameList {
			if v.Gender == gender {
				dataWithGender = append(dataWithGender, v)
			}
		}
	}

	if len(dataWithGender) == 0 {
		dataWithGender = FirstNameList
	}

	if country != "" {
		for _, v := range dataWithGender {
			if v.Country == country {
				data = append(data, v)
			}
		}
	}

	if len(data) == 0 {
		data = FirstNameList
	}

	return data[randInt(data)].Name
}

// Lastname generates a random last name according specified gender and country.
func Lastname(country string) string {
	randSeed()

	var data []Person

	if country != "" {
		for _, v := range LastNameList {
			if v.Country == country {
				data = append(data, v)
			}
		}
	}

	if len(data) == 0 {
		data = LastNameList
	}

	return data[randInt(data)].Name
}

// Fullname generates a random full name according specified gender and country.
func Fullname(gender, country string) string {
	firstName := Firstname(gender, country)
	lastName := Lastname(country)

	return fmt.Sprintf("%s %s", firstName, lastName)
}
