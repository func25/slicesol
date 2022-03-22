package slicesol

import "errors"

var (
	ErrQueueEmpty = errors.New("cannot dequeue because slice is empty")
	ErrStackEmpty = errors.New("cannot pop because slice is empty")
)

// Dequeue pops the first element, simulate a queue
// this function will change the original slice
func (t *Sliol[T]) Dequeue() (T, error) {
	return Dequeue((*[]T)(t))
}

// Dequeue pops the first element, simulate a queue
// this function will change the original slice
func Dequeue[T any](t *[]T) (T, error) {
	n := len(*t)
	if n == 0 {
		return getDefault[T](), ErrQueueEmpty
	}
	res := (*t)[0]
	*t = (*t)[1:n]

	return res, nil
}

// func Push[T any](t *[]T, a T) {
// 	*t = append(*t, a)
// }

// Pop pops the last element, simulate a stack
// this function will change the original slice
func (t *Sliol[T]) Pop() (T, error) {
	return Pop((*[]T)(t))
}

// Pop pops the last element, simulate a stack
// this function will change the original slice
func Pop[T any](t *[]T) (T, error) {
	n := len(*t)
	if n == 0 {
		return getDefault[T](), ErrStackEmpty
	}

	res := (*t)[n-1]
	*t = Remove(*t, n-1)

	return res, nil
}
