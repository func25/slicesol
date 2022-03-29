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

//OptCacheSelected set equal to true to not iterate over selected elements of previous BFS/DFS calls
//if you pass the selected element into the query, it will return nil slice
func OptCacheSelected[T any](cacheSelected bool) SearchOption[T] {
	return func(q *SearchQuery[T]) {
		q.cacheSelected = cacheSelected
	}
}

var justTrue = func(Vector2D) bool { return true }

type IterateOption func(*iterateConfig)

type iterateConfig struct {
	isDFS       bool
	size        Vector2D
	selectFirst func(Vector2D) bool
}

func (c *iterateConfig) applyFrom(opts ...IterateOption) *iterateConfig {
	*c = iterateConfig{
		isDFS:       false,
		size:        Vector2D{},
		selectFirst: justTrue,
	}

	for i := range opts {
		opts[i](c)
	}

	return c
}

func IOptDFS(isDFS bool) IterateOption {
	return func(q *iterateConfig) {
		q.isDFS = isDFS
	}
}

func IOptCustomSize(size Vector2D) IterateOption {
	return func(q *iterateConfig) {
		q.size = size
	}
}

func IOptFirstCondition(fn func(Vector2D) bool) IterateOption {
	return func(q *iterateConfig) {
		q.selectFirst = fn
	}
}

type SearchFunc[T any] func(Vector2D, ...SearchOption[T]) ([]Vector2D, error)
