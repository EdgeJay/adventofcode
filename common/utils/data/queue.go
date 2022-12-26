package data

type Queue[V any] struct {
	list []*V
}

func (q *Queue[V]) In(val *V) {
	q.list = append(q.list, val)
}

func (q *Queue[V]) Out() *V {
	if len(q.list) == 1 {
		val := q.list[0]
		q.list = make([]*V, 0)
		return val
	} else if len(q.list) > 1 {
		val := q.list[0]
		q.list = q.list[1:]
		return val
	}
	return nil
}

func (q *Queue[V]) IsEmpty() bool {
	return len(q.list) == 0
}

func NewQueue[V any]() *Queue[V] {
	return &Queue[V]{
		list: make([]*V, 0),
	}
}
