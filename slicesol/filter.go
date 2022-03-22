package slicesol

// Filter creates
func Filter[T any](a []T, fn func(e T) bool) []T {
	filtered := make([]T, 0)
	for _, e := range a {
		if fn(e) {
			filtered = append(filtered, e)
		}
	}

	return filtered
}
