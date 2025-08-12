package arch

type Queue[T any] []T

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var res T
		return res, false
	}

	item := (*q)[0]
	*q = (*q)[1:]

	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
