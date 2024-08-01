package iterator

// Chain chains multiple iterators into a single iterator
func Chain[T any](iterators ...Iterator[T]) *Chained[T] {
	return &Chained[T]{
		iterators: iterators,
		index:     0,
	}
}

type Chained[T any] struct {
	iterators []Iterator[T]
	index     int
}

func (i *Chained[T]) HasNext() bool {
	if len(i.iterators) == 0 {
		return false
	}
	for !i.iterators[i.index].HasNext() {
		i.index++
		if i.index == len(i.iterators) {
			return false
		}
	}
	return true
}

func (i *Chained[T]) Next() T {
	return i.iterators[i.index].Next()
}
