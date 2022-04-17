package service

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandInt(max, min int64) int64 {
	return rand.Int63n(max-min) + min
}
