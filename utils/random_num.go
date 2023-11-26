package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000000) + 1000000
}
