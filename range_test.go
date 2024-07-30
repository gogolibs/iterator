package iterator_test

import (
	"github.com/gogolibs/iterator"
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	i := iterator.Items(1, 2, 3, 4)
	actual := make([]int, 3)
	iterator.Range[int](i, func(index int, item int) bool {
		if item == 4 {
			return false
		}
		actual[index] = item
		return true
	})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("actual slice %#v is not equal to expected %#v", actual, expected)
	}
}
