package random

import (
	"math/rand"
	"reflect"
	"time"
)

func randSeed() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(v interface{}) int {
	l := reflect.ValueOf(v).Len()

	return rand.Intn(l)
}
