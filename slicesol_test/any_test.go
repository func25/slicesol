package slicesoltest

import (
	"fmt"
	"testing"

	"github.com/func25/slicesol/slicesol"
)

type anyTestTable[T any] struct {
	InpArr  func() []T
	InpFunc func(e T) bool
	Fn      func(arr []T, e T) bool
}

func TestAny(t *testing.T) {
	table := []anyTestTable[int]{
		{
			InpArr: func() []int { return []int{} },
			InpFunc: func(e int) bool {
				return e == 1
			},
			Fn: func(arr []int, e int) bool {
				return e == -1
			},
		},
		{
			InpArr: func() []int { return []int{1} },
			InpFunc: func(e int) bool {
				return e == 1
			},
			Fn: func(arr []int, e int) bool {
				return e != -1 && arr[0] == 1
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			InpFunc: func(e int) bool {
				return e == 3
			},
			Fn: func(arr []int, e int) bool {
				return e != -1 && arr[e] == 3
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} },
			InpFunc: func(e int) bool {
				return e%2 == 0
			},
			Fn: func(arr []int, e int) bool {
				return e != -1 && arr[e]%2 == 0
			},
		},
	}

	t.Run("Any", func(t *testing.T) {
		for k, v := range table {
			arr := v.InpArr()
			if !v.Fn(arr, slicesol.Any(arr, v.InpFunc)) {
				t.Error(fmt.Sprintf("wrong slicesol.Any: %v, test case: %d", arr, k))
			}
		}
	})
}
