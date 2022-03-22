package slicesoltest

import (
	"testing"

	"github.com/func25/slicesol/slicesol"
)

type mapTestTable[T any] struct {
	InpArr  func() []T
	InpFunc func(e T) T
	Fn      func(arr []T, origin []T) bool
}

func TestMap(t *testing.T) {
	table := []mapTestTable[int]{
		{
			InpArr: func() []int { return []int{1} },
			InpFunc: func(e int) int {
				return e * 2
			},
			Fn: func(arr []int, origin []int) bool {
				return len(arr) == len(origin) && arr[0] == 2
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			InpFunc: func(e int) int {
				return e * 2
			},
			Fn: func(arr []int, origin []int) bool {
				if len(arr) != len(origin) {
					return false
				}

				for k := range arr {
					if arr[k] != origin[k]*2 {
						return false
					}
				}

				return true
			},
		},
		{
			InpArr: func() []int { return []int{} },
			Fn: func(arr []int, origin []int) bool {
				return len(arr) == len(origin)
			},
		},
	}

	t.Run("Map", func(t *testing.T) {
		for _, v := range table {
			arr := v.InpArr()
			if !v.Fn(slicesol.Map(arr, v.InpFunc), arr) {
				t.Error("wrong slicesol.Map")
			}
		}
	})
}
