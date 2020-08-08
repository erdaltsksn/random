package random_test

import (
	"log"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/erdaltsksn/random"
)

func TestTC(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TC (Turkish Citizen Number)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := random.TC()
			if !isValidTC(got) {
				t.Error("Got:", got, "It's not a valid TC")
			}
		})
	}
}

func isValidTC(tc string) bool {
	var n [11]int
	for i, v := range strings.Split(tc, "") {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
			return false
		}
		n[i] = num
	}

	if n[0] == 0 {
		return false
	}

	c1 := (7 * (n[0] + n[2] + n[4] + n[6] + n[8])) - (n[1] + n[3] + n[5] + n[7])
	c1 = int(math.Mod(float64(c1), 10.0))
	if n[9] != c1 {
		return false
	}

	c2 := (n[0] + n[1] + n[2] + n[3] + n[4] + n[5] + n[6] + n[7] + n[8] + c1)
	c2 = int(math.Mod(float64(c2), 10.0))
	if n[10] != c2 {
		return false
	}

	return true
}
