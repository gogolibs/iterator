package iterator_test

import (
	"github.com/gogolibs/iterator"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	i := iterator.Items(1, 2, 3)
	if !i.HasNext() {
		t.Fatal("must have first value")
	}
	v1 := i.Next()
	if v1 != 1 {
		t.Fatal("first value must be 1")
	}
	if !i.HasNext() {
		t.Fatal("must have second value")
	}
	v2 := i.Next()
	if v2 != 2 {
		t.Fatal("first value must be 2")
	}
	if !i.HasNext() {
		t.Fatal("must have third value")
	}
	v3 := i.Next()
	if v3 != 3 {
		t.Fatal("first value must be 3")
	}
	if i.HasNext() {
		t.Fatal("must not have more than 3 values")
	}
}

func TestSliceNextEmpty(t *testing.T) {
	expected := "runtime error: index out of range [0] with length 0"
	i := iterator.Items[int]()
	defer func() {
		r := recover()
		err, ok := r.(error)
		if !ok {
			t.Fatalf(`expected to recover from error, got %s instead`, reflect.TypeOf(r))
		}
		actual := err.Error()
		if actual != expected {
			t.Fatalf(`error message "%s" is different from expected "%s"`, expected, actual)
		}
	}()
	i.Next()
}
