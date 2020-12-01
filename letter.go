package random

import (
	"math/rand"
)

// Letter generate a random lowercase or uppercase letter.
func Letter() string {
	randSeed()

	letterCase := []int{'a', 'A'}
	r := rand.Intn(2)

	return string(rune(letterCase[r] + rand.Intn(26)))
}

// LetterLowerCase generate a random lowercase letter.
func LetterLowerCase() string {
	randSeed()

	return string(rune('a' + rand.Intn(26)))
}

// LetterUpperCase generate a random uppercase letter.
func LetterUpperCase() string {
	randSeed()

	return string(rune('A' + rand.Intn(26)))
}
