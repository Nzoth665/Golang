package main

import (
	"fmt"
	"os"

	// "strconv"
	"database/sql"
	"os/user"
	"strings"

	"github.com/bogem/id3v2/v2"
	_ "github.com/mattn/go-sqlite3"
)

const (
	musicAddresC = "C:/Users/@USER@/AppData/Local/Packages/A025C540.Yandex.Music_vfvw9svesycw6/LocalState/Music/80e3b70ca63a2e69a01337e83c068452"
	dbAddresC    = "C:/Users/@USER@/AppData/Local/Packages/A025C540.Yandex.Music_vfvw9svesycw6/LocalState/musicdb_80e3b70ca63a2e69a01337e83c068452.sqlite"
	request      = "SELECT T_Track.Title AS Track_name, T_Album.Title AS Album, T_Album.ArtistsString AS Artist, T_Track.Type FROM T_Album JOIN T_Track JOIN T_TrackAlbum ON T_Track.Id = T_TrackAlbum.TrackId AND T_Album.Id = T_TrackAlbum.AlbumId AND T_Track.Id = $1"
)

type metadata struct {
	TrackName string
	AlbumName string
	Artist    string
	Type      string
	Autor     string
	Num       string
}

func (m *metadata) bookPandler() {
	if m.Type == "audiobook" {
		t := strings.Replace(strings.Replace(m.TrackName, "«", "~", 1), "»", "~", 1)
		q := strings.Split(t, "~")
		num := strings.Split(q[2], " ")[2]
		m.Num = strings.Repeat("0", 2-len(num)) + num
		m.Autor = strings.Replace(q[0], ". ", "", 1)
		m.TrackName = q[1] + ". Часть " + m.Num
		m.AlbumName = q[1]
	}
}

func main() {
	user, err := user.Current()
	usern := strings.Split(user.Username, "\\")[1]
	if err != nil {
		panic(err.Error())
	}
	musicAddres := strings.Replace(musicAddresC, "@USER@", usern, 1)
	dbAddres := strings.Replace(dbAddresC, "@USER@", usern, 1)

	dir, err := os.Open(musicAddres)
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", dbAddres)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	a := id3v2.EncodingUTF16

	m := metadata{}
	for _, file := range files {
		//tag.AddTextFrame(tag.CommonID(""), a, )
		meta := db.QueryRow(request, strings.Replace(file.Name(), ".mp3", "", 1))
		if err = meta.Scan(&m.TrackName, &m.AlbumName, &m.Artist, &m.Type); err != nil {
			panic(err)
		}
		m.bookPandler()

		tag, err := id3v2.Open(musicAddres+"/"+file.Name(), id3v2.Options{Parse: true})
		if err != nil {
			panic(err)
		}

		tag.AddTextFrame(tag.CommonID("Title"), a, m.TrackName)
		tag.AddTextFrame(tag.CommonID("Track number/Position in set"), a, m.Num)
		tag.AddTextFrame(tag.CommonID("Album/Movie/Show title"), a, m.AlbumName)
		tag.AddTextFrame(tag.CommonID("Content type"), a, m.Type)
		tag.AddTextFrame(tag.CommonID("Lead artist/Lead performer/Soloist/Performing group"), a, m.Autor)
		tag.AddTextFrame(tag.CommonID("Band/Orchestra/Accompaniment"), a, m.Artist)
		if err = tag.Save(); err != nil {
			panic(err)
		}
		tag.Close()

	L1:
		err = os.Rename(musicAddres+"/"+file.Name(), "./"+m.AlbumName+"/"+m.TrackName+".mp3")
		if err != nil {
			os.Mkdir(m.AlbumName, 7089)
			fmt.Println(m)
			goto L1
		}
	}
}
