package v1_test

import (
	"testing"

	random "github.com/erdaltsksn/random/v1"
)

var allLetters = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

var lowercaseLettters = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

var uppercaseLetters = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func TestLetter(t *testing.T) {
	type args struct {
		lower bool
		upper bool
	}
	tests := []struct {
		name   string
		args   args
		output []string
	}{
		{"Letter", args{}, allLetters},
		{"Lowercase", args{lower: true, upper: false}, lowercaseLettters},
		{"Only Lowercase", args{lower: true}, lowercaseLettters},
		{"Uppercase", args{lower: false, upper: true}, uppercaseLetters},
		{"Only Uppercase", args{upper: true}, uppercaseLetters},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.Letter(tt.args.lower, tt.args.upper)
			t.Run(tt.name, func(t *testing.T) {
				if !inLetters(got, tt.output) {
					t.Error("Got:", got, ",", tt.output)
				}
			})
		}
	}
}

func inLetters(d string, letters []string) bool {
	for _, v := range letters {
		if v == d {
			return true
		}
	}
	return false
}
