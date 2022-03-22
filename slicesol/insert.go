package slicesol

// Insert inserts elems to a at i (index)
// if i > len(elems) than result will be extended
// so Insert([1,2,3], 4, [5, 6]) will return [1, 2, 3, 0, 5, 6]
func (t Sliol[T]) Insert(i int, elems ...T) Sliol[T] {
	return Insert(t, i, elems...)
}

// Insert inserts elems to a at i (index)
// if i > len(elems) than result will be extended
// so Insert([1,2,3], 4, [5, 6]) will return [1, 2, 3, 0, 5, 6]
func Insert[T any](a []T, i int, elems ...T) []T {
	if len(elems) == 0 || i < 0 {
		return a
	}
	if i > len(a) {
		a = extend(a, i-len(a))
	}

	a = append(a[:i], append(elems, a[i:]...)...)
	return a
}
