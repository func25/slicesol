package slicesoltest

import (
	"fmt"
	"testing"

	"github.com/func25/slicesol/slicesol"
)

type structureTestTable[T any] struct {
	InpArr func() []T
	Fn     func(pop T, err error, arr []T) bool
}

func TestStructure(t *testing.T) {
	popTable := []structureTestTable[int]{
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			Fn: func(pop int, err error, arr []int) bool {
				return err == nil && len(arr) == 4 && pop == 5
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3, 4} },
			Fn: func(pop int, err error, arr []int) bool {
				return err == nil && len(arr) == 3 && pop == 4
			},
		},
		{
			InpArr: func() []int { return []int{1, 2, 3} },
			Fn: func(pop int, err error, arr []int) bool {
				return err == nil && len(arr) == 2 && pop == 3
			},
		},
		{
			InpArr: func() []int { return []int{1, 2} },
			Fn: func(pop int, err error, arr []int) bool {
				return err == nil && len(arr) == 1 && pop == 2
			},
		},
		{
			InpArr: func() []int { return []int{1} },
			Fn: func(pop int, err error, arr []int) bool {
				return err == nil && len(arr) == 0 && pop == 1
			},
		},
		{
			InpArr: func() []int { return []int{} },
			Fn: func(pop int, err error, arr []int) bool {
				return err == slicesol.ErrStackEmpty && pop == 0 && len(arr) == 0
			},
		},
	}

	dequeueTable := []structureTestTable[int]{
		{
			InpArr: func() []int { return []int{1, 2, 3, 4, 5} },
			Fn: func(deq int, err error, arr []int) bool {
				return err == nil && len(arr) == 4 && deq == 1
			},
		},
		{
			InpArr: func() []int { return []int{2, 3, 4, 5} },
			Fn: func(deq int, err error, arr []int) bool {
				return err == nil && len(arr) == 3 && deq == 2
			},
		},
		{
			InpArr: func() []int { return []int{3, 4, 5} },
			Fn: func(deq int, err error, arr []int) bool {
				return err == nil && len(arr) == 2 && deq == 3
			},
		},
		{
			InpArr: func() []int { return []int{4, 5} },
			Fn: func(deq int, err error, arr []int) bool {
				return err == nil && len(arr) == 1 && deq == 4
			},
		},
		{
			InpArr: func() []int { return []int{5} },
			Fn: func(deq int, err error, arr []int) bool {
				return err == nil && len(arr) == 0 && deq == 5
			},
		},
		{
			InpArr: func() []int { return []int{} },
			Fn: func(deq int, err error, arr []int) bool {
				return err == slicesol.ErrQueueEmpty && deq == 0 && len(arr) == 0
			},
		},
	}

	t.Run("Pop", func(t *testing.T) {
		for k, v := range popTable {
			arr := v.InpArr()
			pop, err := slicesol.Pop(&arr)
			if !v.Fn(pop, err, arr) {
				t.Error(fmt.Sprintf("wrong slicesol.Pop: %v, test case: %d", arr, k))
			}
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		for k, v := range dequeueTable {
			arr := v.InpArr()
			deq, err := slicesol.Dequeue(&arr)
			if !v.Fn(deq, err, arr) {
				t.Error(fmt.Sprintf("wrong slicesol.Dequeue: %v, test case: %d", arr, k))
			}
		}
	})
}
