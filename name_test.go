package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

func TestFirstName(t *testing.T) {
	type args struct {
		gender  string
		country string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Empty", args{}},
		{"Male", args{gender: "Male"}},
		{"Female", args{gender: "Female"}},
		{"Invalid Gender", args{gender: "x"}},
		{"Country 1", args{country: "USA"}},
		{"Country 2", args{country: "Turkey"}},
		{"Country Invalid", args{country: "Invalid_Country"}},
		{"Both Gender % Country 1", args{gender: "Male", country: "Turkey"}},
		{"Both Gender % Country 2", args{gender: "Female", country: "Russia"}},
		{"Invalid Gender % Country", args{gender: "Alian", country: "Invalid"}},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.Firstname(tt.args.gender, tt.args.country)
			t.Run(tt.name, func(t *testing.T) {
				if !inNames(random.FirstNameList, got, tt.args.gender, tt.args.country) {
					t.Error("Got:", got, "Want:", random.FirstNameList)
				}
			})
		}
	}
}

func TestLastName(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Empty", args{}},
		{"Country 1", args{country: "USA"}},
		{"Country 2", args{country: "Turkey"}},
		{"Country Invalid", args{country: "Invalid_Country"}},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.Lastname(tt.args.country)
			t.Run(tt.name, func(t *testing.T) {
				if !inNames(random.LastNameList, got, "", tt.args.country) {
					t.Error("Got:", got, "Want:", random.LastNameList)
				}
			})
		}
	}
}

func inNames(list []random.Person, n, gender, country string) bool {
	var dataWithGender []random.Person
	var data []random.Person

	if gender != "" {
		for _, v := range list {
			if v.Gender == gender {
				dataWithGender = append(dataWithGender, v)
			}
		}
	}

	if len(dataWithGender) == 0 {
		dataWithGender = list
	}

	if country != "" {
		for _, v := range dataWithGender {
			if v.Country == country {
				data = append(data, v)
			}
		}
	}

	if len(data) == 0 {
		data = list
	}

	for _, v := range data {
		if v.Name == n {
			return true
		}
	}
	return false
}
