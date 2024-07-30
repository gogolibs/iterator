package iterator

func Items[T any](items ...T) *SliceIterator[T] {
	return Slice(items)
}

func Slice[T any](s []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice: s,
		index: 0,
	}
}

type SliceIterator[T any] struct {
	slice []T
	index int
}

func (i *SliceIterator[T]) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *SliceIterator[T]) Next() T {
	item := i.slice[i.index]
	i.index++
	return item
}
