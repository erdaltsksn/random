package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

func TestCountry(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Countries"},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.Country()
			t.Run(tt.name, func(t *testing.T) {
				if !inCountries(got) {
					t.Error("Got:", got, "Want:", random.CountryList)
				}
			})
		}
	}
}

func inCountries(c string) bool {
	for _, v := range random.CountryList {
		if v == c {
			return true
		}
	}
	return false
}
