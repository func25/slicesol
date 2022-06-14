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

// MapID transforms each element of input slice through fn and return a new transformed slice
func (t Sliol[T]) MapID(fn func(e T, id int) T) Sliol[T] {
	return MapID(t, fn)
}

// MapID transforms each element of input slice through fn and return a new transformed slice
func MapID[T, K any](a []T, fn func(e T, id int) K) []K {
	res := make([]K, len(a))
	for k, v := range a {
		res[k] = fn(v, k)
	}

	return res
}

// MapWithError transforms each element of input slice through fn and return a new transformed slice
//It were interupted if any error occured
func (t Sliol[T]) MapErr(fn func(e T) (T, error)) (Sliol[T], error) {
	return MapErr(t, fn)
}

// MapWithError transforms each element of input slice through fn and return a new transformed slice
//It were interupted if any error occured
func MapErr[T, K any](a []T, fn func(e T) (K, error)) ([]K, error) {
	var err error
	res := make([]K, len(a))
	for k, v := range a {
		res[k], err = fn(v)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
