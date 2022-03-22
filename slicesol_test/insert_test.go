package slicesoltest

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/func25/slicesol/slicesol"
)

type insertTestTable[T any] struct {
	InpArr    func() []T
	InpID     int
	Insertion []T
	Fn        func(arr []T) bool
}

func TestInsert(t *testing.T) {
	table := []insertTestTable[int]{
		{
			InpArr:    func() []int { return []int{1} },
			InpID:     0,
			Insertion: []int{2, 3, 4, 5},
			Fn: func(arr []int) bool {
				return reflect.DeepEqual(arr, []int{2, 3, 4, 5, 1})
			},
		},
		{
			InpArr:    func() []int { return []int{} },
			Insertion: []int{1},
			InpID:     0,
			Fn: func(arr []int) bool {
				return len(arr) == 1 && arr[0] == 1
			},
		},
		{
			InpArr:    func() []int { return []int{1, 2, 3, 4, 5} },
			InpID:     0,
			Insertion: []int{},
			Fn: func(arr []int) bool {
				return len(arr) == 5
			},
		},
		{
			InpArr:    func() []int { return []int{1, 2, 3} },
			Insertion: []int{5},
			InpID:     4,
			Fn: func(arr []int) bool {
				return len(arr) == 5 && reflect.DeepEqual(arr, []int{1, 2, 3, 0, 5})
			},
		},
	}

	t.Run("Insert", func(t *testing.T) {
		for k, v := range table {
			arr := v.InpArr()
			if !v.Fn(slicesol.Insert(arr, v.InpID, v.Insertion...)) {
				t.Error(fmt.Sprintf("wrong slicesol.Insert: %v, test case %d", v.InpArr(), k))
			}
		}
	})
}
