package structers

// Стопка
type Stack[T any] struct {
	massive []T
	len     int
}

// Возращает верхний элемент стопки
func (s *Stack[T]) Get() T {
	return s.massive[s.len-1]
}

// Кладёт элемент на верх стопки
func (s *Stack[T]) Put(a T) {
	if s.len == len(s.massive) {
		s.len += 1
		s.massive = append(s.massive, a)
	} else {
		s.len += 1
		s.massive[s.len-1] = a
	}
}

// Удаляет элемент с вершины стопки
func (s *Stack[T]) Pop() {
	if s.len-1 >= 0 {
		s.len -= 1
		//s.massive = s.massive[:s.len]
	}
}

// Удаляет все элементы в стопке
func (s *Stack[T]) Clean() {
	s.len = 0
	s.massive = []T{}
}

func (s *Stack[T]) Len() int {
	return s.len
}
