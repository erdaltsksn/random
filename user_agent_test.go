package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

func TestUserAgent(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"User Agents"},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.UserAgent()
			t.Run(tt.name, func(t *testing.T) {
				if !inUserAgents(got) {
					t.Error("Got:", got, "Want:", random.UserAgentList)
				}
			})
		}
	}
}

func inUserAgents(c string) bool {
	for _, v := range random.UserAgentList {
		if v == c {
			return true
		}
	}
	return false
}
