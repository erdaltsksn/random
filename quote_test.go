package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

func TestQuote(t *testing.T) {
	type args struct {
		author string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Empty", args{}},
		{"Non-author", args{author: "No born author"}},
		{"With Author", args{author: "Oscar Wilde"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				got := random.Quote(tt.args.author)
				t.Run(tt.name, func(t *testing.T) {
					if !inQuotes(got.Author) {
						t.Error("Got:", got, "Want:", "random.QuoteList")
					}
				})
			}
		})
	}
}

func inQuotes(author string) bool {
	var data []random.QuoteStr

	if author != "" {
		for _, v := range random.QuoteList {
			if v.Author == author {
				data = append(data, v)
			}
		}
	}

	if len(data) == 0 {
		data = random.QuoteList
	}

	for _, v := range data {
		if v.Author == author {
			return true
		}
	}
	return false
}
