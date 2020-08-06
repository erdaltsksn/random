package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Person is a struct hold data.
type Person struct {
	Name    string
	Gender  string
	Country string
}

// Firstname generates a random first name according specified gender and country.
func Firstname(gender, country string) string {
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

	rand.Seed(time.Now().Unix())

	return data[rand.Intn(len(data))].Name
}

// Lastname generates a random last name according specified gender and country.
func Lastname(country string) string {
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

	rand.Seed(time.Now().Unix())

	return data[rand.Intn(len(data))].Name
}

// Fullname generates a random full name according specified gender and country.
func Fullname(gender, country string) string {
	firstName := Firstname(gender, country)
	lastName := Lastname(country)

	return fmt.Sprintf("%s %s", firstName, lastName)
}
