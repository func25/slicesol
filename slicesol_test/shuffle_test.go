package slicesoltest

import (
	"fmt"
	"testing"

	"github.com/func25/slicesol/slicesol"
)

type shuffleTestTable[T any] struct {
	InpArr func() []T
	Fn     func(arr []T, origin []T) bool
}

func TestShuffle(t *testing.T) {
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
				fmt.Println(arr)
				for k := range arr {
					if arr[k] != origin[k] {
						return true
					}
				}

				return false
			},
		},
		{
			InpArr: func() []int { return []int{} },
			Fn: func(arr []int, origin []int) bool {
				return len(arr) == 0
			},
		},
	}

	t.Run("Shuffle", func(t *testing.T) {
		for _, v := range table {
			arr := v.InpArr()
			if !v.Fn(slicesol.Shuffle(arr), arr) {
				t.Error("wrong slicesol.Shuffle")
			}
		}
	})

	t.Run("Shuffle with Ref", func(t *testing.T) {
		for _, v := range table {
			arr := v.InpArr()
			slicesol.ShuffleRef(arr)
			if !v.Fn(slicesol.ShuffleRef(arr), v.InpArr()) {
				t.Error("wrong slicesol.ShuffleRef")
			}
		}
	})
}
