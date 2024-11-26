package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

type MaG struct {
	forks       [][2]int
	coordinates [2]int
	matrix      *[][]int
}

func (g *MaG) move() {
	m := *g.matrix
L1:
	ch := check(g.coordinates, m)
	if len(ch) > 0 {
		if len(ch) > 1 {
			g.forks = append(g.forks, g.coordinates)
		}
		d := ch[rand.Intn(len(ch))]
		switch d {
		case "↑":
			m[g.coordinates[0]-1][g.coordinates[1]] = 0
			g.coordinates = [2]int{g.coordinates[0] - 2, g.coordinates[1]}
		case "→":
			m[g.coordinates[0]][g.coordinates[1]+1] = 0
			g.coordinates = [2]int{g.coordinates[0], g.coordinates[1] + 2}
		case "↓":
			m[g.coordinates[0]+1][g.coordinates[1]] = 0
			g.coordinates = [2]int{g.coordinates[0] + 2, g.coordinates[1]}
		case "←":
			m[g.coordinates[0]][g.coordinates[1]-1] = 0
			g.coordinates = [2]int{g.coordinates[0], g.coordinates[1] - 2}
		}
	} else if len(ch) == 0 {
		g.coordinates = g.forks[len(g.forks)-1]
		g.forks = g.forks[:len(g.forks)-1]
		goto L1
	}
	m[g.coordinates[0]][g.coordinates[1]] = 0
	*g.matrix = m
}

// проверка на наличие не посещённых клеток вокруг некоторой клетки
/*
   s - кордината некоторой клетки (верхняя левая клетка имеет координаты 0, 0)
   m - сама матрица либиринта
*/
func check(s [2]int, m [][]int) []string {
	var b []string
	if s[0]-2 > 0 {
		if m[s[0]-2][s[1]] == -1 {
			b = append(b, "↑")
		}
	}
	if s[1]+2 < len(m[0]) {
		if m[s[0]][s[1]+2] == -1 {
			b = append(b, "→")
		}
	}
	if s[0]+2 < len(m) {
		if m[s[0]+2][s[1]] == -1 {
			b = append(b, "↓")
		}
	}
	if s[1]-2 > 0 {
		if m[s[0]][s[1]-2] == -1 {
			b = append(b, "←")
		}
	}
	return b
}

func mazeGenerator(height int, width int, kq int) {
	k := float64(kq) / 100.0
	var matrix [][]int
	var s1 []int
	var s2 []int
	mag := MaG{coordinates: [2]int{1, 1}, forks: [][2]int{}, matrix: &matrix}

	// Созданние 2-ух мерного массива
	for i := 0; i < height; i++ {
		s1 = append(s1, 1)
		if i%2 != 0 {
			s2 = append(s2, -1)
		} else {
			s2 = append(s2, 1)
		}
	}
	for i := 0; i < width; i++ {
		if i%2 != 0 {
			matrix = append(matrix, append([]int(nil), s2...))
		} else {
			matrix = append(matrix, append([]int(nil), s1...))
		}
	}
	matrix[1][1] = 1

	// Обработка нужных переменных
	var ndc [][2]int
	var passed int
	for y := range matrix {
		for x := range matrix[y] {
			if (y%2 != 0) && (x%2 != 0) {
				passed += 1
			}
			if ((x+y)%2 != 0) && (x*y != 0) && (x != width-1) && (y != height-1) {
				ndc = append(ndc, [2]int{y, x})
			}
		}
	}

	// Генерация лабиринта
	for passed != 1 {
		mag.move()
		passed -= 1
	}
	q := rand.Perm(len(ndc))
	w := int(float64(len(ndc)) * k)
	fmt.Println(w)
	for i := 0; i < w; i++ {
		matrix[ndc[q[i]][0]][ndc[q[i]][1]] = 0
	}

	// Преобразование 2-ух мерного массива в изображение
	matrix[1][1] = 2
	matrix[width-2][height-2] = 3
	img := image.NewRGBA(image.Rect(0, 0, height, width))
	for y := range matrix {
		for x := range matrix[y] {
			switch matrix[y][x] {
			case 0:
				img.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			case 1:
				img.Set(x, y, color.RGBA{A: 255})
			case 2:
				img.Set(x, y, color.RGBA{B: 255, A: 255})
			case 3:
				img.Set(x, y, color.RGBA{R: 255, A: 255})
			}
		}
	}

	// сохранение изображения
	file, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}
