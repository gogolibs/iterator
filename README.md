# iterator

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

## Iterator implementation for slices

```go

package main

import (
	"fmt"
	"github.com/gogolibs/iterator"
)

func main() {
	i := iterator.FromItems(1, 2, 3)
	fmt.Println(i.HasNext())               // true
	fmt.Println(i.Next())                  // 1
	fmt.Println(i.HasNext())               // true
	fmt.Println(i.Next())                  // 2
	fmt.Println(i.HasNext())               // true
	fmt.Println(i.Next())                  // 3
	fmt.Println(i.HasNext())               // false
	fmt.Println(i.Next())                  // panics
	i = iterator.FromSlice([]int{1, 2, 3}) // a different way to create iterator from slice
	s := iterator.ToSlice[int](i, 3)       // need to specify size of the slice here
	fmt.Println(s)                         // prints 1, 2, 3
}

```

## ForEach function

```go
package main

import (
	"fmt"
	"github.com/gogolibs/iterator"
)

func main() {
	i := iterator.FromItems(1, 2, 3)
	hasFinished := iterator.ForEach(i, func(item int) bool {
		if item == 2 {
			return false // stop iteration if the item is 2
		}
		fmt.Println(item) // will only print 1
		return true       // continue iteration
	})
	fmt.Println(hasFinished) // prints false because we stopped at 2
}

```

## Chaining multiple iterators

```go

package main

import (
	"fmt"
	"github.com/gogolibs/iterator"
)

func main() {
	i1 := iterator.FromItems(1, 2)
	i2 := iterator.FromItems(3, 4)
	chained := iterator.Chain[int](i1, i2)
	fmt.Println(iterator.ToSlice[int](chained, 4)) // prints 1, 2, 3, 4
}

```

## Transforming an iterator

```go

package main

import (
	"fmt"
	"github.com/gogolibs/iterator"
)

func main() {
	i := iterator.FromItems(1, 2, 3)
	t := iterator.Transform(i, func(item int) string {
		return fmt.Sprintf("user%d", item)
	})
	fmt.Println(iterator.ToSlice[string](t, 3)) // prints user1, user2, user3
}

```