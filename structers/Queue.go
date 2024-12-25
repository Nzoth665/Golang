package structers

// Очередь
type Queue[T any] struct {
	massive     []T
	firselement int
	len         int
	//capacity    int
}

func CreateQueue[T any]() Queue[T] { return Queue[T]{[]T{}, 0, 0} }

// Возращает первый элемент очереди
func (q *Queue[T]) Get() T {
	return q.massive[q.firselement]
}

// Кладёт элемент в конец очереди
func (q *Queue[T]) Put(a T) {
	q.len++
	q.massive = append(q.massive, a)
}

// Удаляет элемент с начала очереди
func (q *Queue[T]) Pop() {
	q.firselement++
	q.len--
}

// Удаляет неиспользуемые элементы из очереди
func (q *Queue[T]) Clean() {
	q.massive = q.massive[q.firselement : q.len+q.firselement]
	q.len = q.firselement
	q.firselement = 0
}

// Удаляет все элементы из очереди
func (q *Queue[T]) CleanAll() {
	q.massive = []T{}
	q.len = 0
	q.firselement = 0
}

// Возращает длину очереди
func (q *Queue[T]) Len() int { return q.len }
