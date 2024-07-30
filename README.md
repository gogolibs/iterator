# iterator #

[![GoDoc](https://godoc.org/github.com/gogolibs/iterator?status.svg)](https://pkg.go.dev/github.com/gogolibs/iterator)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogolibs/iterator?style=flat)](https://goreportcard.com/report/github.com/gogolibs/iterator)
[![CI](https://github.com/gogolibs/iterator/actions/workflows/ci.yml/badge.svg)](https://github.com/gogolibs/iterator/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/gogolibs/iterator/graph/badge.svg?token=6Z4I4YR035)](https://codecov.io/gh/gogolibs/iterator)

**iterator** provides a generic `Iterator[T]` interface:

```go
package iterator

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
```

and also contains an implementation for slices:

```go

package example

import (
	"fmt"
	"github.com/gogolibs/iterator"
)

func example() {
	i := iterator.Items(1, 2, 3)
	fmt.Println(i.HasNext()) // true
	fmt.Println(i.Next())    // 1
	fmt.Println(i.HasNext()) // true
	fmt.Println(i.Next())    // 2
	fmt.Println(i.HasNext()) // true
	fmt.Println(i.Next())    // 3
	fmt.Println(i.HasNext()) // false
	fmt.Println(i.Next())    // panics
}

```

and range function to easily iterate through iterator items:
```go

package example

import (
	"fmt"
	"github.com/gogolibs/iterator"
)

func example() {
	i := iterator.Items(1, 2, 3, 4)
	iterator.Range[int](i, func(index int, item int) bool {
        if item == 4 {
			return false // bool return argument simulates break behaviour -- return false to stop iteration
		}
		fmt.Println(item)
		return true
	}) // prints 1, 2, 3
}

```