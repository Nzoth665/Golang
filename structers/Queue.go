package structers

type Query struct {
	massive     []any
	firselement int
	Len         int
}

func (q *Query) Get() any {
	return q.massive[q.firselement]
}

func (q *Query) Put(a any) {
	q.Len++
	q.massive = append(q.massive, a)
}

func (q *Query) Pop() {
	q.firselement++
	q.Len--
}

func (q *Query) Clear() {
	q.massive = q.massive[q.firselement : q.Len+q.firselement]
}
