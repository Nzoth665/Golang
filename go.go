package main

import (
	"fmt"
)

type Chain[T any] struct {
	a     T
	chain *Chain[T]
}

func Atoc[T any](a []T, b bool) []Chain[T] {
	m := []Chain[T]{}
	for i, e := range a {
		if i != 0 {
			m = append(m, Chain[T]{e, nil})
			m[i-1].chain = &m[i]
		} else {
			m = append(m, Chain[T]{e, nil})
		}
	}
	return m
}

func (ch *Chain[T]) Next() {
	if ch.chain != nil {
		ch = ch.chain
	}
}

/*type Chain struct {
	a int
	chain *Chain
}

func Atoc(a []int, b bool) (ch Chain) {
	m := []Chain{}
}

func (ch *Chain) Next() {
	if ch.chain != nil {
		ch = ch.chain
	}
}*/

func main() {
	a := []int{1, 2, 3, 4, 5}
	m := Atoc(a, false)
	c := m[0]
	fmt.Println(m)
	fmt.Println(c)
	//c.Next()
	fmt.Println(*c.chain)
	fmt.Println(m[1].chain)

	fmt.Println()

	c1 := Chain[int]{1, nil}
	c2 := Chain[int]{2, &c1}
	c3 := Chain[int]{3, &c2}
	c4 := Chain[int]{4, &c3}
	c5 := Chain[int]{5, &c4}
	fmt.Println(c5)
	fmt.Println(*c5.chain)
	fmt.Println(*(*c5.chain).chain)
}
