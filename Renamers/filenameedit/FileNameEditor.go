package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
)

func input(reader *bufio.Reader) string {
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.Trim(text, "\n\t\r")
}

func dir_init(dirname string) (dir *os.File, files []os.FileInfo) {
	dir, err := os.Open(dirname)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	files, err = dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	files = slices.DeleteFunc(files, func(f os.FileInfo) bool {
		return (f.Name() == "FileNameEditor.go" || f.Name() == "renamer.exe")
	})
	//fmt.Println(files)
	return
}

func main() {
	dir, files := dir_init(".")

	reader := bufio.NewReader(os.Stdin)
	v := false
	f := func() {
		if !v {
			fmt.Println("This command invalid")
		}
		v = false
	}

	fmt.Print("Comands:\n1) delete\n\t1.1) all\n\t1.2) words\n\t\t1.2.1) first\n\t\t1.2.2) end\n\t\t1.2.3) center\n2) replace\n3) content\n4) update\n5) move\n\t5.1) to\n\t5.2) now\n\t5.3) out\n6) end or exit\n\n")

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("This command invalid")
		}
	}()
	for {
		request := strings.Split(input(reader), " ")

		switch request[0] {
		case "delete":
			switch request[1] {
			case "all":
				text := input(reader)
				for _, e := range files {
					if e.IsDir() {
						continue
					}
					os.Rename(path.Join(dir.Name(), e.Name()), path.Join(dir.Name(), strings.Replace(e.Name(), text, "", 1)))
				}
				v = true
			case "words":
				i := 0
				fmt.Scan(&i)
				switch request[2] {
				case "first":
					for _, e := range files {
						if e.IsDir() {
							continue
						}
						os.Rename(path.Join(dir.Name(), e.Name()), path.Join(dir.Name(), strings.Join(strings.Split(e.Name(), " ")[i:], " ")))
					}
				case "end":
					for _, e := range files {
						if e.IsDir() {
							continue
						}
						m := strings.Split(e.Name(), " ")
						os.Rename(path.Join(dir.Name(), e.Name()), path.Join(dir.Name(), strings.Join(m[:len(m)-i], " ")))
					}
				case "center":
					j := 0
					fmt.Scan(&j)
					for _, e := range files {
						if e.IsDir() {
							continue
						}
						m := strings.Split(e.Name(), " ")
						os.Rename(path.Join(dir.Name(), e.Name()), path.Join(dir.Name(), strings.Join(append(m[:i], m[j:]...), " ")))
					}
				default:
					f()
				}
				v = true
			default:
				f()
				continue
			}
			fmt.Println("\"delete\": OK")

		case "move":
			switch request[1] {
			case "to":
				b := true
				i := 0
				for i, e := range files {
					if e.IsDir() {
						fmt.Print(i, ") ", e.Name(), "\n")
						b = false
					}
				}
				if b {
					fmt.Println("No dirs!")
					goto End
				}
				v = true
				fmt.Scan(&i)
				dir, files = dir_init(path.Join(dir.Name(), files[i].Name()))
			End:
				fmt.Println(dir.Name())
			case "now":
				fmt.Println(dir.Name())
			case "out":
				dir, files = dir_init(path.Dir(dir.Name()))
				fmt.Println(dir.Name())
			default:
				f()
			}

		case "replace":
			text1 := input(reader)
			text2 := input(reader)
			for _, e := range files {
				if e.IsDir() {
					continue
				}
				os.Rename(path.Join(dir.Name(), e.Name()), path.Join(dir.Name(), strings.Replace(e.Name(), text1, text2, 1)))
			}
			v = true
			fmt.Println("\"replace\": OK")

		case "content":
			for i, e := range files {
				fmt.Printf("%d) %s", i, e.Name())
				if e.IsDir() {
					fmt.Println(" (is dir)")
				} else {
					fmt.Println("")
				}
			}

		case "update":
			dir, files = dir_init(dir.Name())
			fmt.Println("\"update\": OK")

		case "end", "exit":
			return
		default:
			f()
		}
	}
}
