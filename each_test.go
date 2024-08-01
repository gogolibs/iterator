package iterator_test

import (
	"github.com/gogolibs/iterator"
	"reflect"
	"testing"
)

func TestForEach(t *testing.T) {
	testCases := map[string]struct {
		input           []int
		f               func([]int, int) ([]int, bool)
		expected        []int
		expectExhausted bool
	}{
		"1-2-3-4, exhausted": {
			input: []int{1, 2, 3, 4},
			f: func(s []int, i int) ([]int, bool) {
				return append(s, i), true
			},
			expected:        []int{1, 2, 3, 4},
			expectExhausted: true,
		},
		"1-2-3-4, not exhausted": {
			input: []int{1, 2, 3, 4},
			f: func(s []int, i int) ([]int, bool) {
				if i == 4 {
					return s, false
				}
				return append(s, i), true
			},
			expected:        []int{1, 2, 3},
			expectExhausted: false,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			i := iterator.FromSlice(testCase.input)
			var slice []int
			isExhausted := iterator.ForEach[int](i, func(item int) bool {
				var proceed bool
				slice, proceed = testCase.f(slice, item)
				return proceed
			})
			if !reflect.DeepEqual(testCase.expected, slice) {
				t.Fatalf("actual slice %#v is not equal to expected %#v", slice, testCase.expected)
			}
			if isExhausted != testCase.expectExhausted {
				if isExhausted {
					t.Fatal("iterator is exhausted but expected not to be")
				} else {
					t.Fatal("iterator is not exhausted but expected to be")
				}
			}
		})
	}
}
