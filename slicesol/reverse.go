package slicesol

// Reverse will CHANGE the original slice and its element order also
func (t *Sliol[T]) Reverse(i int) Sliol[T] {
	return Reverse(*t)
}

// Reverse will CHANGE the original slice and its element order also
func Reverse[T any](a []T) []T {
	x1, x2 := 0, len(a)-1
	for x1 < x2 {
		a[x1], a[x2] = a[x2], a[x1]
		x1++
		x2--
	}

	return a
}
