package slicesoltest

import (
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
