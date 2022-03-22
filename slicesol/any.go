package slicesol

// Any returns index of element that meet fn condition, or -1 if not
func (t Sliol[T]) Any(fn func(T) bool) int {
	return Any(t, fn)
}

// Any returns index of element that meet fn condition, or -1 if not
func Any[T any](a []T, fn func(T) bool) int {
	for k, v := range a {
		if fn(v) {
			return k
		}
	}

	return -1
}

// IsAny will return true if any element meets fn condition, or -1 if not
func (t Sliol[T]) IsAny(fn func(T) bool) bool {
	return IsAny(t, fn)
}

// IsAny will return true if any element meets fn condition, or -1 if not
func IsAny[T any](a []T, fn func(T) bool) bool {
	for _, v := range a {
		if fn(v) {
			return true
		}
	}

	return false
}
