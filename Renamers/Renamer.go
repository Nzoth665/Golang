package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Открываем текущую директорию
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	// Получаем список файлов и папок
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := ""
	fmt.Scanln(&s)
	s = strings.ReplaceAll(s, "_", " ")

	// Выводим имена файлов и папок
	for i, file := range files {
		if file.Name() != "Renamer.exe" {
			//fmt.Println(i, file.Name(), s)
			err = os.Rename(file.Name(), s+" Глава "+strconv.Itoa(i+1)+".mp3")
			if err != nil {
				panic(err)
			}
		}
	}
}
