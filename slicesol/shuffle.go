package slicesol

import (
	"math/rand"
	"time"
)

// Shuffle will return a NEW slice as shuffled version of input slice
func (t Sliol[T]) Shuffle() Sliol[T] {
	return Shuffle(t)
}

// Shuffle will return shuffled version of input slice
func Shuffle[T any](slice []T) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}
