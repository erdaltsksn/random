package v1

import (
	"math/rand"
	"time"
)

// Digit generate a random digit
func Digit() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(10)
}
