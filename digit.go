package random

import (
	"math/rand"
)

// Digit generate a random digit.
func Digit() int {
	randSeed()

	return rand.Intn(10)
}
