package random

import (
	"math/rand"
	"time"
)

func randSeed() {
	rand.Seed(time.Now().UnixNano())
}
