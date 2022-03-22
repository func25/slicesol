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
