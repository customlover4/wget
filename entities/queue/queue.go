package queue

type el struct {
	Data any
	next *el
}

type Queue struct {
	first *el
	last  *el
	len   int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Add(data any) {
	q.len++
	el := &el{data, nil}
	if q.first == nil {
		q.first = el
		q.last = el
	}

	q.last.next = el
	q.last = el
}

func (q *Queue) Pop() any {
	if q.first == nil || q.last == nil {
		return nil
	}
	q.len--
	if q.first == q.last {
		el := q.first
		q.first = nil
		q.last = nil
		return el.Data
	}

	el := q.first
	q.first = q.first.next

	return el.Data
}

func (q *Queue) Empty() bool {
	if q.first == nil && q.last == nil {
		return true
	}
	return false
}

func (q *Queue) Len() int {
	return q.len
}
