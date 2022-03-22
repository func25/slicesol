package slicesoltest

import (
	"testing"

	"github.com/func25/slicesol/slicesol"
)

func TestDuplicate(t *testing.T) {
	table := []shuffleTestTable[int]{
		{
			InpArr: func() []int { return []int{1} },
			Fn: func(arr []int, origin []int) bool {
				return len(arr) == len(origin) && arr[0] == 1
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			Fn: func(arr []int, origin []int) bool {
				if len(arr) != len(origin) {
					return false
				}

				for k := range arr {
					if arr[k] != origin[k] {
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

	t.Run("Duplicate", func(t *testing.T) {
		for _, v := range table {
			arr := v.InpArr()
			if !v.Fn(slicesol.Duplicate(arr), arr) {
				t.Error("wrong slicesol.Duplicate")
			}
		}
	})
}
