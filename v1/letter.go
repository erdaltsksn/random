package v1

import (
	"math/rand"
	"time"
)

// Letter generate a random letter
func Letter(lower bool, upper bool) string {
	rand.Seed(time.Now().UnixNano())

	if lower {
		return string('a' + rand.Intn(26))
	}

	if upper {
		return string('A' + rand.Intn(26))
	}

	c := []int{'a', 'A'}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)
	return string(c[r] + rand.Intn(26))
}
