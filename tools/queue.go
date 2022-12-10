package tools

type Queue[T any] struct {
	enq *Stack[T]
	deq *Stack[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		enq: NewStack[T](),
		deq: NewStack[T](),
	}
}

func (q *Queue[T]) Enqueue(elem T) {
	q.enq.Push(elem)
}

func (q *Queue[T]) Dequeue() T {
	if q.deq.Length() == 0 {
		if q.enq.Length() == 0 {
			var null T
			return null
		}
		for range q.enq.List() {
			q.deq.Push(q.enq.Pop())
		}
	}
	return q.deq.Pop()
}

func (q *Queue[T]) List() []T {
	return append(q.enq.List(), q.deq.Reverse().List()...)
}
