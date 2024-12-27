package structers

type Pointer[T any] struct {
	massive *[]T
	element int
}

func (p *Pointer[T]) Next() {
	if p.element < len(*p.massive) {
		p.element++
	}
}

func (p *Pointer[T]) Get() T {
	m := *p.massive
	return m[p.element]
}

func (p *Pointer[T]) ILE() bool { return len(*p.massive) == p.element }

// ///////////////////////////////////////////////////////////////////////////////////////
type Pointers[T any] []Pointer[T]

// int|int16|int32|int64|int8|float32|float64|complex128|complex64
type PointersNum[T int | int16 | int32 | int64 | int8 | float32 | float64] []Pointer[T]

func (a PointersNum[T]) Len() int      { return len(a) }
func (a PointersNum[T]) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PointersNum[T]) Less(i, j int) bool {
	return (*a[i].massive)[a[i].element] < (*a[j].massive)[a[j].element]
}
