package iterator_test

import (
	"github.com/gogolibs/iterator"
	"reflect"
	"strconv"
	"testing"
)

func TestTransform_IntToString(t *testing.T) {
	i := iterator.FromItems(1, 2, 3)
	actual := iterator.ToSlice[string](iterator.Transform[int, string](i, strconv.Itoa), 3)
	expected := []string{"1", "2", "3"}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("actual slice %#v is not equal to expected %#v", actual, expected)
	}
}
