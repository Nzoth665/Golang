package structers

type Stack struct {
	Massive []any
	len     int
	//capacity int
}

func (s *Stack) Get() any {
	return s.Massive[s.len-1]
}

func (s *Stack) Put(a any) {
	s.len += 1
	s.Massive = append(s.Massive, a)
}

func (s *Stack) Pop() {
	if s.len-1 >= 0 {
		s.len -= 1
		s.Massive = s.Massive[:s.len-1]
	}
}
