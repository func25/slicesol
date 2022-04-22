package slicesol

import "github.com/func25/mathfunc"

// Map transforms each element of input slice through fn and return a new transformed slice
func (t Sliol[T]) Rand() T {
	return Rand(t)
}

// Map transforms each element of input slice through fn and return a new transformed slice
func Rand[T any](a []T) T {
	id, _ := mathfunc.Random0ToInt(len(a))
	return a[id]
}
