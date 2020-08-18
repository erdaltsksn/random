package random

import (
	"math/rand"
	"time"
)

// QuoteStr is a struct hold data.
type QuoteStr struct {
	Text   string
	Author string
}

// Quote generates a random quote. You can limit it by specifying an author.
func Quote(author string) QuoteStr {
	var data []QuoteStr

	if author != "" {
		for _, v := range QuoteList {
			if v.Author == author {
				data = append(data, v)
			}
		}
	}

	if len(data) == 0 {
		data = QuoteList
	}

	rand.Seed(time.Now().Unix())

	quote := data[rand.Intn(len(data))]

	return quote
}
