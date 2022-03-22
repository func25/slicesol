package slicesol

// RemoveUnor will CHANGE the original slice
// usage:
//		arr = RemoveUnor(arr, i)
// DO NOT:
// 		other = RemoveUnor(arr, i)
func RemoveUnor[T any](a []T, i int) []T {
	n := len(a)

	if i >= len(a) || i < 0 {
		return a
	}

	a[i] = a[n-1]
	a[n-1] = getDefault[T]()
	a = a[:n-1]

	return a
}

// Remove will CHANGE the original slice
// usage:
//		arr = Remove(arr, i)
// DO NOT:
// 		other = Remove(arr, i)
func Remove[T any](a []T, i int) []T {
	n := len(a)

	if i >= n || i < 0 {
		return a
	}

	copy(a[i:], a[i+1:])
	a[n-1] = getDefault[T]()
	a = a[:n-1]

	return a
}
