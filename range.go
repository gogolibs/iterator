package iterator

func Range[T any](iterator Iterator[T], f func(index int, item T)) {
	index := 0
	for iterator.HasNext() {
		f(index, iterator.Next())
		index++
	}
}
