package slicesoltest

import (
	"testing"

	"github.com/func25/slicesol/slicesol"
)

type filterTestTable[T any] struct {
	InpArr  func() []T
	InpFunc func(e T) bool
	Fn      func(arr []T) bool
}

func TestFilter(t *testing.T) {
	table := []filterTestTable[int]{
		{
			InpArr: func() []int { return []int{1} },
			InpFunc: func(e int) bool {
				return e%2 == 0
			},
			Fn: func(arr []int) bool {
				return len(arr) == 0
			},
		},
		{
			InpArr: func() []int { return []int{1, 3} },
			InpFunc: func(e int) bool {
				return e%2 == 0
			},
			Fn: func(arr []int) bool {
				return len(arr) == 0
			},
		},
		{
			InpArr: func() []int { return []int{2} },
			InpFunc: func(e int) bool {
				return e%2 == 0
			},
			Fn: func(arr []int) bool {
				return len(arr) == 1 && arr[0] == 2
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			InpFunc: func(e int) bool {
				return e%2 == 0
			},
			Fn: func(arr []int) bool {
				return len(arr) == 2 && arr[0] == 2 && arr[1] == 4
			},
		},
		{
			InpArr: func() []int { return []int{} },
			Fn: func(arr []int) bool {
				return len(arr) == 0
			},
		},
	}

	t.Run("Filter", func(t *testing.T) {
		for _, v := range table {
			arr := v.InpArr()
			if !v.Fn(slicesol.Filter(arr, v.InpFunc)) {
				t.Error("wrong slicesol.Filter")
			}
		}
	})
}
