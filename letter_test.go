package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

func TestLetter(t *testing.T) {
	var allLetters = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	tests := []struct {
		name string
	}{
		{"All Letters"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				got := random.Letter()
				t.Run(tt.name, func(t *testing.T) {
					if !inLetters(got, allLetters) {
						t.Error("Got:", got, "Want:", allLetters)
					}
				})
			}
		})
	}
}

func TestLetterLowerCase(t *testing.T) {
	var lowercaseLetters = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	tests := []struct {
		name string
	}{
		{"Lowercase Letters"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				got := random.LetterLowerCase()
				t.Run(tt.name, func(t *testing.T) {
					if !inLetters(got, lowercaseLetters) {
						t.Error("Got:", got, "Want:", lowercaseLetters)
					}
				})
			}
		})
	}
}

func TestLetterUpperCase(t *testing.T) {
	var uppercaseLetters = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	tests := []struct {
		name string
	}{
		{"Uppercase Letters"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				got := random.LetterUpperCase()
				t.Run(tt.name, func(t *testing.T) {
					if !inLetters(got, uppercaseLetters) {
						t.Error("Got:", got, "Want:", uppercaseLetters)
					}
				})
			}
		})
	}
}

func inLetters(l string, letters []string) bool {
	for _, v := range letters {
		if v == l {
			return true
		}
	}
	return false
}
