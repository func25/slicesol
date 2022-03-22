package slicesol

// Any returns index of element that meet fn condition, or -1 if not
func Any[T any](a []T, fn func(T) bool) int {
	for k, v := range a {
		if fn(v) {
			return k
		}
	}

	return -1
}

// IsAny answers the question that "does any element meet fn condition?"
func IsAny[T any](a []T, fn func(T) bool) bool {
	for _, v := range a {
		if fn(v) {
			return true
		}
	}

	return false
}
