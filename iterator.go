package iterator

// Iterator is an interface for the iterator pattern
// https://en.wikipedia.org/wiki/Iterator_pattern.
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
