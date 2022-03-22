package slicesol

import (
	"math/rand"
	"time"
)

// Shuffle will return a NEW slice as shuffled version of input slice
func Shuffle[T any](slice []T) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	res := Duplicate(slice)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { res[i], res[j] = res[j], res[i] })
	return res
}

// ShuffleRef changes input slice and returns itself also
// no allocate new slice
func ShuffleRef[T any](slice []T) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}
