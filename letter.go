package random

import (
	"math/rand"
	"time"
)

// Letter generate a random lowercase or uppercase letter
func Letter() string {
	letterCase := []int{'a', 'A'}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)

	return string(letterCase[r] + rand.Intn(26))
}

// LetterLowerCase generate a random lowercase letter
func LetterLowerCase() string {
	rand.Seed(time.Now().UnixNano())

	return string('a' + rand.Intn(26))
}

// LetterUpperCase generate a random uppercase letter
func LetterUpperCase() string {
	rand.Seed(time.Now().UnixNano())

	return string('A' + rand.Intn(26))
}
