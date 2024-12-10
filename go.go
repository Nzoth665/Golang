package main

import (
	. "fmt"
	"strings"
)

func solution(students []map[string]string) []string {
	m := []string{}
	for _, e := range students {
		s := strings.ToLower(strings.Replace(e["studentName"], " ", "-", 1))
		m = append(m, s)
	}
	return m
}

func main() {
	var l, d, k int
	Scan(&l, &d, &k)
	c := l / d
	Print(min(c, k-1))
}
