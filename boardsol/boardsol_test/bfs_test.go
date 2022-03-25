package boardsoltest

import (
	"fmt"
	"testing"

	"github.com/func25/slicesol/boardsol"
	"github.com/func25/slicesol/slicesol"
)

type vector2D boardsol.Vector2D

func TestBFS(t *testing.T) {
	b := [][]int{
		{1, 1, 3, 3, 3},
		{2, 1, 1, 4, 3},
		{2, 2, 1, 4, 3},
		{2, 2, 1, 1, 5},
		{5, 5, 5, 4, 5},
		{1, 2, 3, 4, 5},
	}

	q := boardsol.BFSQuery[int]{
		Face: b,
		SelectCondition: func(cur, next boardsol.Vector2D) bool {
			return next.X >= 0 && next.X < len(b) &&
				next.Y >= 0 && next.Y < len(b[next.X]) &&
				b[cur.X][cur.Y] == b[next.X][next.Y]
		},
	}

	largestGroup := []boardsol.Vector2D{}
	largestValue := -1
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			group, err := q.BFS(boardsol.Vector2D{X: i, Y: j})
			if err != nil {
				t.Fatal(err)
			}
			if len(group) > len(largestGroup) {
				largestGroup = group
				largestValue = b[group[0].X][group[0].Y]
			}
		}
	}

	if largestValue != 1 {
		t.Error(`wrong "largest" value`)
	}

	if !slicesol.Equal(largestGroup, []boardsol.Vector2D{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 2}, {3, 2}}) {
		t.Error(`wrong positions`)
	}

	fmt.Println(`"lagest" value:`, largestValue)
	fmt.Println("positions:", largestGroup)
}
