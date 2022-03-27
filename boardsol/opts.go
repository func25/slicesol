package boardsol

type BFSOption[T any] func(*SearchQuery[T])

func (q *SearchQuery[T]) ApplyOpts(opts ...BFSOption[T]) {
	for i := range opts {
		opts[i](q)
	}
}

//OptDirections custom selecting element directions for search query
func OptDirections[T any](dirs []Vector2D) BFSOption[T] {
	return func(q *SearchQuery[T]) {
		q.directions = dirs
	}
}

//OptLimit limit the number of selected elements of each call
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

//OptKeepExists set true to not iterate over selected elements of previous BFS/DFS calls
//if you pass the selected element into the query, it will return nil slice
func OptKeepExists[T any](keepExists bool) BFSOption[T] {
	return func(q *SearchQuery[T]) {
		q.keepExists = keepExists
	}
}
