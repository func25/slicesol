package slicesol

// Filter creates new slice contains elements that meet fn condition
func (t Sliol[T]) Filter(fn func(e T) bool) Sliol[T] {
	return Filter(t, fn)
}

// Filter creates new slice contains elements that meet fn condition
func Filter[T any](a []T, fn func(e T) bool) []T {
	filtered := make([]T, 0)
	for _, e := range a {
		if fn(e) {
			filtered = append(filtered, e)
		}
	}

	return filtered
}
