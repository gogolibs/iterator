package iterator

// ForEach calls function f for each item in the iterator.
// Function f is required to return a flag signalling whether
// iteration should be continued or not. ForEach itself returns
// a flag signalling whether the iteration proceeded to the end (true)
// or not (false).
func ForEach[T any](iterator Iterator[T], f func(item T) bool) bool {
	for iterator.HasNext() {
		if !f(iterator.Next()) {
			return false
		}
	}
	return true
}
