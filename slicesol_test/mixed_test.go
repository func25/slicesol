package slicesoltest

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/func25/slicesol/slicesol"
)

func TestPipe(t *testing.T) {
	t.Run("Mixed", func(t *testing.T) {
		arr := slicesol.Sliol[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		newSlice := arr.Filter(func(e int) bool { return e%2 == 0 }).
			Map(func(e int) int { return e * 2 })
		if !reflect.DeepEqual(([]int)(newSlice), []int{4, 8, 12, 16, 20}) {
			t.Error("wrong mixed", newSlice)
		}
	})
}

func TestRemoveReadme(t *testing.T) {
	arr := slicesol.Remove([]int{1, 2, 3, 4, 5}, 0)
	fmt.Println(arr)
}

func TestInsertReadme(t *testing.T) {
	arr := slicesol.Insert([]int{1, 4, 5}, 1, 2, 3)
	fmt.Println(arr)
}

func TestPopReadme(t *testing.T) {
	popArr := &[]int{1, 2, 3, 4, 5}
	p, err := slicesol.Pop(popArr)
	fmt.Println(p, err)
	fmt.Println(popArr)

	deqArr := &[]int{1, 2, 3, 4, 5}
	d, err := slicesol.Dequeue(deqArr)
	fmt.Println(d, err)
	fmt.Println(deqArr)
}
