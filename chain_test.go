package iterator_test

import (
	"github.com/gogolibs/iterator"
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {
	testCases := map[string]struct {
		input    [][]int
		expected []int
	}{
		"empty": {
			input:    [][]int{},
			expected: []int{},
		},
		"single 1-2-3": {
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		"couple 1-2-3": {
			input: [][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			expected: []int{1, 2, 3, 1, 2, 3},
		},
		"couple 1-2-3 with empty in-between": {
			input: [][]int{
				{1, 2, 3},
				{},
				{1, 2, 3},
			},
			expected: []int{1, 2, 3, 1, 2, 3},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			totalLen := 0
			iterators := make([]iterator.Iterator[int], len(testCase.input))
			for index, slice := range testCase.input {
				totalLen += len(slice)
				iterators[index] = iterator.FromSlice(slice)
			}
			slice := iterator.ToSlice[int](iterator.Chain(iterators...), totalLen)
			if !reflect.DeepEqual(testCase.expected, slice) {
				t.Fatalf("actual slice %#v is not equal to expected %#v", slice, testCase.expected)
			}
		})
	}
}
