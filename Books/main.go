package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bogem/id3v2/v2"
)

const musicAddres = "C:/Users/Max/Downloads/Аркан"

func main() {
	dir, err := os.Open(musicAddres)
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	l := len(strconv.Itoa(len(files)))

	name := ""
	fmt.Scan(&name)
	name = strings.ReplaceAll(name, "_", " ")
	os.Mkdir(name, 0765)
	autor := ""
	fmt.Scan(&autor)
	autor = strings.ReplaceAll(autor, "_", " ")

	for i, file := range files {
		tag, err := id3v2.Open(musicAddres+"/"+file.Name(), id3v2.Options{Parse: true})
		if err != nil {
			panic(err)
		}
		q := strconv.Itoa(i + 1)
		//tag.AddTextFrame(tag.CommonID(""), a, )
		a := id3v2.EncodingUTF16
		tag.AddTextFrame(tag.CommonID("Title"), a, name)
		tag.AddTextFrame(tag.CommonID("Track number/Position in set"), a, strings.Repeat("0", l-len(q))+q)
		tag.AddTextFrame(tag.CommonID("Album/Movie/Show title"), a, name)
		tag.AddTextFrame(tag.CommonID("Content type"), a, "Audiobook")
		tag.AddTextFrame(tag.CommonID("Lead artist/Lead performer/Soloist/Performing group"), a, autor)
		tag.AddTextFrame(tag.CommonID("Attached picture"), a, "icon.jpg")
		if err = tag.Save(); err != nil {
			panic(err)
		}
		tag.Close()

		err = os.Rename(musicAddres+"/"+file.Name(), "./"+name+"/"+name+" Глава "+strings.Repeat("0", l-len(q))+q+".mp3")
		if err != nil {
			panic(err)
		}
	}
}
