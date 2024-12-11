package structers

type Queue struct {
	massive     []any
	firselement int
	Len         int
}

func (q *Queue) Get() any {
	return q.massive[q.firselement]
}

func (q *Queue) Put(a any) {
	q.Len++
	q.massive = append(q.massive, a)
}

func (q *Queue) Pop() {
	q.firselement++
	q.Len--
}

func (q *Queue) Clean() {
	q.massive = q.massive[q.firselement : q.Len+q.firselement]
}

func (q *Queue) CleanAll() {
	q.massive = []any{}
}
