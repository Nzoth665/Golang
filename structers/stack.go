package structers

// Стопка
type Stack struct {
	massive []any
	len     int
}

// Возращает верхний элемент стопки
func (s *Stack) Get() any {
	return s.massive[s.len-1]
}

// Кладёт элемент на верх стопки
func (s *Stack) Put(a any) {
	s.len += 1
	s.massive = append(s.massive, a)
}

// Удаляет элемент с вершины стопки
func (s *Stack) Pop() {
	if s.len-1 >= 0 {
		s.len -= 1
		s.massive = s.massive[:s.len]
	}
}

// Удаляет все элементы в стопке
func (s *Stack) Clean() {
	s.len = 0
	s.massive = []any{}
}
