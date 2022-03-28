package boardsol

type SearchOption[T any] func(*SearchQuery[T])

func (q *SearchQuery[T]) ApplyOpts(opts ...SearchOption[T]) *SearchQuery[T] {
	for i := range opts {
		opts[i](q)
	}

	return q
}

//OptDirections custom selecting element directions for search query
func OptDirections[T any](dirs []Vector2D) SearchOption[T] {
	return func(q *SearchQuery[T]) {
		q.directions = dirs
	}
}

//OptLimit limit the number of selected elements of each call
func OptLimit[T any](limit int) SearchOption[T] {
	return func(q *SearchQuery[T]) {
		q.limit = limit
	}
}

func OptGenDirections[T any](fn func() []Vector2D) SearchOption[T] {
	return func(q *SearchQuery[T]) {
		q.genDirections = fn
	}
}

//OptCacheSelectedElements set equal to true to not iterate over selected elements of previous BFS/DFS calls
//if you pass the selected element into the query, it will return nil slice
func OptCacheSelectedElements[T any](cacheElements bool) SearchOption[T] {
	return func(q *SearchQuery[T]) {
		q.cacheElements = cacheElements
	}
}

type IterateOption func(*iterateConfig)

type iterateConfig struct {
	isBFS bool
	size  Vector2D
}

func (c *iterateConfig) applyOpts(opts ...IterateOption) *iterateConfig {
	for i := range opts {
		opts[i](c)
	}

	return c
}

func IOptBFS(isBFS bool) IterateOption {
	return func(q *iterateConfig) {
		q.isBFS = isBFS
	}
}

func IOptCustomSize(size Vector2D) IterateOption {
	return func(q *iterateConfig) {
		q.size = size
	}
}

type SearchFunc[T any] func(Vector2D, ...SearchOption[T]) ([]Vector2D, error)
