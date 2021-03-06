package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

var digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestDigit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Digit"},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.Digit()
			t.Run(tt.name, func(t *testing.T) {
				if !inDigits(got) {
					t.Error("Got:", got, "Want:", digits)
				}
			})
		}
	}
}

func inDigits(d int) bool {
	for _, v := range digits {
		if v == d {
			return true
		}
	}
	return false
}
