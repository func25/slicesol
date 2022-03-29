package boardsol

import (
	"errors"

	"github.com/func25/slicesol/slicesol"
)

var ErrEmptyBoard = errors.New("cannot bfs board with width == 0")
var emptyStruct = struct{}{}

//SearchQuery is used for implementing bf-search, df-search
type SearchQuery[T any] struct {
	SelectCondition func(cur Vector2D, next Vector2D) bool // condition to add next element into queue
	Width           int                                    // width of the board

	// config

	// directions: Which directions should the query check?
	// by default, there are 4: up, down, left and right
	directions []Vector2D

	// genDirections: if you want to generate new set of directions each time searching at specific element
	genDirections func() []Vector2D
	limit         int  // limit the number of the selected elements
	cacheSelected bool // cache the selected elements of previous calls

	// state
	exists map[int]struct{} // to not iterate over already selected cells
	res    []Vector2D       // results
}

//BFS returns selected elements that meet specific condition
//if keepExists == true and the pos was already selected then BFS will return nil slice
func (q *SearchQuery[T]) BFS(pos Vector2D, opts ...SearchOption[T]) ([]Vector2D, error) {
	for i := range opts {
		opts[i](q)
	}

	// if cacheElements and the pos was already selected in previous calls
	// return nil slice
	if q.cacheSelected && q.exists != nil {
		if _, exist := q.exists[pos.To1D(q.Width)]; exist {
			return nil, nil
		}
	}

	if err := q.validate(pos); err != nil {
		return nil, err
	}

	queue := slicesol.Sliol[Vector2D]{Vector2D{pos.X, pos.Y}}

	for len(queue) > 0 {
		cur, err := queue.Dequeue()
		if err != nil {
			return nil, err
		}

		for _, dir := range _DIRECTIONS {

			next := cur.Plus(dir)
			id1D := next.To1D(q.Width)

			if _, exist := q.exists[id1D]; !exist && q.SelectCondition(cur, next) {
				queue = append(queue, next)
				q.exists[id1D] = emptyStruct

				q.res = append(q.res, next)

				if q.limit != 0 && len(q.res) > q.limit {
					return q.res, nil
				}
			}
		}
	}

	return q.res, nil
}

func (q *SearchQuery[T]) validate(pos Vector2D) error {
	if q.directions == nil {
		if q.genDirections != nil {
			q.directions = q.genDirections()
		} else {
			q.directions = _DIRECTIONS
		}
	}

	if q.Width == 0 {
		return ErrEmptyBoard
	}

	q.res = []Vector2D{{pos.X, pos.Y}}

	if !q.cacheSelected || q.exists == nil {
		q.exists = map[int]struct{}{}
	}

	q.exists[pos.To1D(q.Width)] = emptyStruct

	return nil
}

//Reset uses to reset the query state like result, selected elements of previous calls
func (q *SearchQuery[T]) Reset() {
	q.res = make([]Vector2D, 0)
	q.exists = make(map[int]struct{})
}

func (q *SearchQuery[T]) Iterate(board [][]T, fn func(group []Vector2D), opts ...IterateOption) error {
	cfg := (&iterateConfig{}).applyFrom(opts...)

	// custom size
	customSize := false
	width := len(board)
	if cfg.size != (Vector2D{}) {
		customSize = true
		width = cfg.size.X
	}

	// custom query
	query := q.BFS
	if cfg.isDFS {
		query = q.DFS
	}

	for i := 0; i < width; i++ {
		lenY := len(board[i])
		if customSize && lenY > cfg.size.Y {
			lenY = cfg.size.Y
		}

		for j := 0; j < lenY; j++ {
			pos := Vector2D{X: i, Y: j}
			if !cfg.selectFirst(pos) {
				continue
			}

			if group, err := query(pos); err != nil {
				return err
			} else {
				fn(group)
			}
		}
	}

	return nil
}
