package slicesol

// Duplicate creates a new slice base on input slice,
// be careful if any element of the slice is pointer/ map/ slice...
// or struct containing pointer/ map/ slice...
func Duplicate[T any](slice []T) []T {
	clone := make([]T, len(slice))
	copy(clone, slice)
	return clone
}
