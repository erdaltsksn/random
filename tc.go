package random

import (
	"fmt"
	"math"
	"math/rand"
)

// TC generates a random Turkish Citizen number.
func TC() string {
	randSeed()

	var (
		d1, d2, d3, d4, d5, d6, d7, d8, d9, c1, c2 int
	)

	// The first number cannot be 0
	d1 = rand.Intn(9) + 1
	d2 = rand.Intn(10)
	d3 = rand.Intn(10)
	d4 = rand.Intn(10)
	d5 = rand.Intn(10)
	d6 = rand.Intn(10)
	d7 = rand.Intn(10)
	d8 = rand.Intn(10)
	d9 = rand.Intn(10)

	c1 = (7 * (d1 + d3 + d5 + d7 + d9)) - (d2 + d4 + d6 + d8)
	c1 = int(math.Mod(float64(c1), 10.0))
	c2 = (d1 + d2 + d3 + d4 + d5 + d6 + d7 + d8 + d9 + c1)
	c2 = int(math.Mod(float64(c2), 10.0))

	tc := fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d", d1, d2, d3, d4, d5, d6, d7, d8, d9, c1, c2)
	return tc
}
