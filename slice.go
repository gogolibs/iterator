package iterator

// FromItems is a synactic sugar for FromSlice function
// for easier inline Slice iterator instantiation.
func FromItems[T any](items ...T) *Slice[T] {
	return FromSlice(items)
}

// FromSlice creates a Slice iterator from a slice.
func FromSlice[T any](s []T) *Slice[T] {
	return &Slice[T]{
		slice: s,
		index: 0,
	}
}

// Slice is an implementation of Iterator for slices.
type Slice[T any] struct {
	slice []T
	index int
}

func (i *Slice[T]) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *Slice[T]) Next() T {
	item := i.slice[i.index]
	i.index++
	return item
}

// ToSlice simply exhausts the iterator into a slice.
// The initial size must be specified as Iterator does not
// have any knowledge of the size of the thing it iterates over.
func ToSlice[T any](iterator Iterator[T], size int) []T {
	slice := make([]T, 0, size)
	ForEach(iterator, func(item T) bool {
		slice = append(slice, item)
		return true
	})
	return slice
}
