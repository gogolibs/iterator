package iterator

// Range function calls f function for each item.
// Boolean return value of the f controls whether iteration should continue or not.
func Range[T any](iterator Iterator[T], f func(index int, item T) bool) {
	index := 0
	for iterator.HasNext() && f(index, iterator.Next()) {
		index++
	}
}
