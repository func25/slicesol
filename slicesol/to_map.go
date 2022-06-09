package slicesol

func ToMap[K comparable, T, V any](t []T, f func(e T) (K, V)) map[K]V {
	m := make(map[K]V, len(t))

	for i := range t {
		k, v := f(t[i])
		m[k] = v
	}

	return m
}

func ToSimpleMap[K comparable, V any](t []V, f func(e V) (K, V)) map[K]V {
	m := make(map[K]V, len(t))

	for i := range t {
		k, v := f(t[i])
		m[k] = v
	}

	return m
}
