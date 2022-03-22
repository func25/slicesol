package slicesoltest

import (
	"testing"

	"github.com/func25/slicesol/slicesol"
)

type RemoveTestTable[T any] struct {
	InpArr func() []T
	InpID  int
	Fn     func(arr []T) bool
}

func TestRemove(t *testing.T) {
	orderTable := []RemoveTestTable[int]{
		{
			InpID:  0,
			InpArr: func() []int { return []int{1} },
			Fn: func(arr []int) bool {
				return len(arr) == 0
			},
		},
		{
			InpID:  0,
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			Fn: func(arr []int) bool {
				return len(arr) == 4 && arr[0] == 2
			},
		},
		{
			InpID:  0,
			InpArr: func() []int { return []int{} },
			Fn: func(arr []int) bool {
				return len(arr) == 0
			},
		},
	}

	unorderTable := []RemoveTestTable[int]{
		{
			InpID:  0,
			InpArr: func() []int { return []int{1} },
			Fn: func(arr []int) bool {
				return len(arr) == 0
			},
		},
		{
			InpID:  0,
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			Fn: func(arr []int) bool {
				return len(arr) == 4
			},
		},
		{
			InpID:  0,
			InpArr: func() []int { return []int{} },
			Fn: func(arr []int) bool {
				return len(arr) == 0
			},
		},
	}

	t.Run("Remove with maintaining order", func(t *testing.T) {
		for _, v := range orderTable {
			arr := v.InpArr()
			arr = slicesol.Remove(arr, v.InpID)
			if !v.Fn(arr) {
				t.Error("wrong slicesol.Remove")
			}
			// fmt.Println(arr)
		}
	})

	t.Run("Remove without order", func(t *testing.T) {
		for _, v := range unorderTable {
			arr := v.InpArr()
			arr = slicesol.Remove(arr, v.InpID)
			if !v.Fn(arr) {
				t.Error("wrong slicesol.Remove")
			}
			// fmt.Println(arr)
		}
	})
}
