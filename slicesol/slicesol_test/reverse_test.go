package slicesoltest

import (
	"testing"

	"github.com/func25/slicesol/slicesol"
)

func TestRerverse(t *testing.T) {
	table := []shuffleTestTable[int]{
		{
			InpArr: func() []int { return []int{1} },
			Fn: func(arr []int, origin []int) bool {
				return len(arr) == 1 && arr[0] == 1
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			Fn: func(arr []int, origin []int) bool {
				if len(arr) != len(origin) {
					return false
				}

				for i, j := 0, len(origin)-1; i < len(arr); i, j = i+1, j-1 {
					if arr[i] != origin[j] {
						return false
					}
				}

				return true
			},
		},
		{
			InpArr: func() []int { return []int{} },
			Fn: func(arr []int, origin []int) bool {
				return len(arr) == 0
			},
		},
	}

	t.Run("Reverse", func(t *testing.T) {
		for _, v := range table {
			arr := v.InpArr()
			origin := v.InpArr()
			if !v.Fn(slicesol.Reverse(arr), origin) {
				t.Error("wrong slicesol.Reverse")
			}
		}
	})
}
