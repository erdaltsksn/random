package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

func TestState(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"States"},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.State()
			t.Run(tt.name, func(t *testing.T) {
				if !inStates(got) {
					t.Error("Got:", got, "Want:", random.StateList)
				}
			})
		}
	}
}

func inStates(c string) bool {
	for _, v := range random.StateList {
		if v == c {
			return true
		}
	}
	return false
}
