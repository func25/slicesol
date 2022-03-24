package boardsol

import (
	"errors"

	"github.com/func25/slicesol/slicesol"
)

var ErrEmptyBoard = errors.New("cannot bfs board with width == 0")

type BFSQuery[T any] struct {
	IsNextValidFunc func(cur T, next T) bool // condition to add next element into queue
	SelectHook      func(Vector2D)           // hook on adding element[x][y] into queue
	Face            [][]T
	Pos             Vector2D

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

func BFSAny[T any](q BFSQuery[T], opts ...BFSOption[T]) error {
	for i := range opts {
		opts[i](&q)
	}

	if q.directions == nil {
		q.directions = _DIRECTIONS
	}

	if q.SelectHook == nil {
		q.SelectHook = func(vd Vector2D) {}
	}

	v0 := Vector2D{}
	if q.size == v0 {
		if len(q.Face) == 0 {
			return ErrEmptyBoard
		}
		q.size.X, q.size.Y = len(q.Face), len(q.Face[0])

		for i := 0; i < q.size.X; i++ {
			if len(q.Face[i]) < q.size.Y {
				q.size.Y = len(q.Face[i])
			}
		}
	}

	queue := slicesol.Sliol[Vector2D]{Vector2D{q.Pos.X, q.Pos.Y}}
	exists := map[int]bool{(q.Pos.Y*q.size.X + q.Pos.X): true}

	for len(queue) > 0 {
		cur, err := queue.Dequeue()
		if err != nil {
			return err
		}

		curValue := q.Face[cur.X][cur.Y]

		for _, dir := range _DIRECTIONS {

			next := cur.plus(dir)
			id1D := next.to1D(q.size.X)

			if !exists[id1D] &&
				q.IsNextValidFunc(curValue, q.Face[next.X][next.Y]) {

				queue = append(queue)
				exists[id1D] = true
				q.SelectHook(next)
			}
		}
	}

	return nil
}
