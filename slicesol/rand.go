package slicesol

import "github.com/func25/mathfunc"

func (t Sliol[T]) Rand() T {
	return Rand(t)
}

func (t Sliol[T]) RandFirst(count int) T {
	return RandFirst(t, count)
}

func Rand[T any](a []T) T {
	id, _ := mathfunc.Random0ToInt(len(a))
	return a[id]
}

func RandFirst[T any](a []T, count int) T {
	if count > len(a) {
		count = len(a)
	}
	id, _ := mathfunc.Random0ToInt(count)
	return a[id]
}
