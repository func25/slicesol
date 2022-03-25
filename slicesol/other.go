package slicesol

func getDefault[T any]() T {
	var def T
	return def
}

func extend[T any](a []T, n int) []T {
	for i := 0; i < n; i++ {
		a = append(a, getDefault[T]())
	}

	return a
}

func Equal[T comparable](a Sliol[T], b Sliol[T]) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}
