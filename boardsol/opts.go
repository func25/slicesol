package boardsol

type BFSOption[T any] func(*SearchQuery[T])

func (q *SearchQuery[T]) ApplyOpts(opts ...BFSOption[T]) {
	for i := range opts {
		opts[i](q)
	}
}

func OptDirections[T any](dirs []Vector2D) BFSOption[T] {
	return func(q *SearchQuery[T]) {
		q.directions = dirs
	}
}

func OptLimit[T any](limit int) BFSOption[T] {
	return func(q *SearchQuery[T]) {
		q.limit = limit
	}
}

func OptGenDirections[T any](fn func() []Vector2D) BFSOption[T] {
	return func(q *SearchQuery[T]) {
		q.genDirections = fn
	}
}

func OptKeepExists[T any](keepExists bool) BFSOption[T] {
	return func(q *SearchQuery[T]) {
		q.keepExists = keepExists
	}
}
