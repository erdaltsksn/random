package random

// QuoteStr is a struct hold data.
type QuoteStr struct {
	Text   string
	Author string
}

// Quote generates a random quote. You can limit it by specifying an author.
func Quote(author string) QuoteStr {
	randSeed()

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

	quote := data[randInt(data)]

	return quote
}
