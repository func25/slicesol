package boardsol

import (
	"errors"

	"github.com/func25/slicesol/slicesol"
)

var ErrEmptyBoard = errors.New("cannot bfs board with width == 0")

type BFSQuery[T any] struct {
	SelectCondition func(cur Vector2D, next Vector2D) bool // condition to add next element into queue
	Face            [][]T

	size       Vector2D
	directions []Vector2D
}

type BFSOption[T any] func(*BFSQuery[T])

func OptSize[T any](width, height int) BFSOption[T] {
	return func(q *BFSQuery[T]) {
		q.size = Vector2D{
			X: width,
			Y: height,
		}
	}
}

func OptDirections[T any](dirs []Vector2D) BFSOption[T] {
	return func(q *BFSQuery[T]) {
		q.directions = dirs
	}
}

func (q *BFSQuery[T]) BFS(pos Vector2D, opts ...BFSOption[T]) ([]Vector2D, error) {
	for i := range opts {
		opts[i](q)
	}

	if q.directions == nil {
		q.directions = _DIRECTIONS
	}

	v0 := Vector2D{}
	if q.size == v0 {
		if len(q.Face) == 0 {
			return nil, ErrEmptyBoard
		}
		q.size.X, q.size.Y = len(q.Face), len(q.Face[0])

		for i := 0; i < q.size.X; i++ {
			if len(q.Face[i]) < q.size.Y {
				q.size.Y = len(q.Face[i])
			}
		}
	}

	queue := slicesol.Sliol[Vector2D]{Vector2D{pos.X, pos.Y}}
	exists := map[int]bool{(pos.Y*q.size.X + pos.X): true}
	res := make([]Vector2D, 0)

	for len(queue) > 0 {
		cur, err := queue.Dequeue()
		if err != nil {
			return nil, err
		}

		for _, dir := range _DIRECTIONS {

			next := cur.plus(dir)
			id1D := next.to1D(q.size.X)

			if !exists[id1D] &&
				q.SelectCondition(cur, next) {

				queue = append(queue, next)
				exists[id1D] = true

				res = append(res, cur)
			}
		}
	}

	return res, nil
}
