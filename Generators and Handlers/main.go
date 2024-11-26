package main

import "fmt"

func main() {
	// Ввод значений высоты, ширины и КДС (Количество Дыр в Стенах)
	var height, width, kq int
	print("Введите ширину, высоту и КДС\n")
	fmt.Scan(&height, &width, &kq)
	mazeGenerator(height, width, kq)
}
