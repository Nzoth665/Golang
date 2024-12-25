package structers

type Pointer[T any] struct {
	massive *[]T
	element int
}

func (p *Pointer[T]) Next() {
	if p.element+1 < len(*p.massive) {
		p.element++
	}
}

func (p *Pointer[T]) Nexti(i int) {
	if p.element+i < len(*p.massive) {
		p.element += i
	} else if p.element < len(*p.massive) && len(*p.massive) < p.element+i {
		p.element = len(*p.massive) - 1
	}
}

func (p *Pointer[T]) Get() T {
	m := *p.massive
	return m[p.element]
}

func (p *Pointer[T]) ILE() bool { return len(*p.massive)-1 == p.element }

type Pointers[T any] []Pointer[T]
