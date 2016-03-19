package ds

import "container/list"

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

func (q *Queue) Enqueue(v int) {
	q.l.PushFront(v)
}

func (q *Queue) Dequeue() int {
	e := q.l.Back()
	v := e.Value.(int)
	q.l.Remove(e)
	return int(v)
}

func (q *Queue) Len() int {
	return q.l.Len()
}
