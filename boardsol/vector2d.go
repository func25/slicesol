package boardsol

type Vector2D struct {
	X, Y int
}

func (pos Vector2D) isBound(size Vector2D) bool {
	return pos.X >= 0 && pos.X < size.X && pos.Y >= 0 && pos.Y < size.Y
}

func (pos Vector2D) plus(other Vector2D) Vector2D {
	return Vector2D{
		X: pos.X + other.X,
		Y: pos.Y + other.Y,
	}
}

func (pos Vector2D) to1D(width int) int {
	return pos.Y*width + pos.X
}

var _DIRECTIONS = []Vector2D{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}
