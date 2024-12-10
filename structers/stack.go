package structers

type Stack struct {
	massive []any
	Len     int
}

func (s *Stack) Get() any {
	return s.massive[s.Len-1]
}

func (s *Stack) Put(a any) {
	s.Len += 1
	s.massive = append(s.massive, a)
}

func (s *Stack) Pop() {
	if s.Len-1 >= 0 {
		s.Len -= 1
		s.massive = s.massive[:s.Len-1]
	}
}

func (s *Stack) Clean() {
	s.Len = 0
	s.massive = []any{}
}
