package slicesol

// Map transforms each element of input slice through fn and return a new transformed slice
func (t Sliol[T]) Map(fn func(e T) T) Sliol[T] {
	return Map(t, fn)
}

// Map transforms each element of input slice through fn and return a new transformed slice
func Map[T, K any](a []T, fn func(e T) K) []K {
	res := make([]K, len(a))
	for k, v := range a {
		res[k] = fn(v)
	}

	return res
}
