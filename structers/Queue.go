package structers

// Очередь
type Queue struct {
	massive     []any
	firselement int
	len         int
	//capacity    int
}

// func CreateQueue(capacity int) Queue { return Queue{[]any{}, 0, 0, capacity} }
func CreateQueue() Queue { return Queue{[]any{}, 0, 0} }

// Возращает первый элемент очереди
func (q *Queue) Get() any {
	return q.massive[q.firselement]
}

// Кладёт элемент в конец очереди
func (q *Queue) Put(a any) {
	q.len++
	q.massive = append(q.massive, a)
}

// Удаляет элемент с начала очереди
func (q *Queue) Pop() {
	q.firselement++
	q.len--
}

// Удаляет неиспользуемые элементы из очереди
func (q *Queue) Clean() {
	q.massive = q.massive[q.firselement : q.len+q.firselement]
	q.len = q.firselement
	q.firselement = 0
}

// Удаляет все элементы из очереди
func (q *Queue) CleanAll() {
	q.massive = []any{}
	q.len = 0
	q.firselement = 0
}

// Возращает длину очереди
func (q *Queue) Len() int { return q.len }
