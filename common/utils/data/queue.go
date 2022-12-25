package data

type Queue[V any] struct {
	list []*V
}

func (q *Queue[V]) In(val *V) {
	q.list = append(q.list, val)
}

func (q *Queue[V]) Out() *V {
	if len(q.list) == 1 {
		q.list = make([]*V, 0)
		return q.list[0]
	} else if len(q.list) > 1 {
		val := q.list[0]
		q.list = q.list[1:]
		return val
	}
	return nil
}
