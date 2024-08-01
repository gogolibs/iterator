package iterator

// Transform transforms Iterator[I] into Iterator[O] by calling func(I) O for each iteration.
// Useful when transformation of the data iterated by Iterator is required,
// but it's not desired or not possible to put the entirety of the data into memory.
func Transform[I, O any](iterator Iterator[I], transform func(I) O) Transformed[I, O] {
	return Transformed[I, O]{
		iterator:  iterator,
		transform: transform,
	}
}

// Transformed iterator is returned by the Transform function. See its docs for more information.
type Transformed[I, O any] struct {
	iterator  Iterator[I]
	transform func(I) O
}

func (i Transformed[I, O]) HasNext() bool {
	return i.iterator.HasNext()
}

func (i Transformed[I, O]) Next() O {
	return i.transform(i.iterator.Next())
}
