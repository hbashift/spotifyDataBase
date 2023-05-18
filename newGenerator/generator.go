package newGenerator

import (
	"db_seminar/structs"
	"encoding/csv"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func ParseDB(db *sqlx.DB) ([]structs.Musician, []structs.Podcast, []structs.Track, []structs.User, []structs.Album) {

	var musicians []structs.Musician
	var podcasts []structs.Podcast
	var tracks []structs.Track
	var users []structs.User
	var albums []structs.Album

	err := db.Select(&musicians, "SELECT musician_id FROM musician")

	if err != nil {
		fmt.Println("musician")
		fmt.Println(err)
	}

	err = db.Select(&podcasts, "SELECT podcast_id FROM podcast")

	if err != nil {
		fmt.Println("podcast")
		fmt.Println(err)
	}

	err = db.Select(&tracks, "SELECT track_id FROM track")

	if err != nil {
		fmt.Println("track")
		fmt.Println(err)
	}

	err = db.Select(&users, "SELECT user_id FROM \"user\"")

	if err != nil {
		fmt.Println("user")
		fmt.Println(err)
	}

	err = db.Select(&albums, "SELECT album_id FROM album")

	if err != nil {
		fmt.Println("album")
		fmt.Println(err)
	}

	return musicians, podcasts, tracks, users, albums
}

func CheckUsers(user1, user2 *structs.User, users []structs.User) (*structs.User, *structs.User) {
	checker := false

	for checker != true {
		if user1.UserId != user2.UserId {
			checker = true
			return user1, user2
		}

		*user1 = users[rand.Intn(len(users))]
		*user2 = users[rand.Intn(len(users))]

		CheckUsers(user1, user2, users)
	}

	return nil, nil
}

func CreateTables(db *sqlx.DB) {
	path, _ := ioutil.ReadFile("newSQLScripts/spotify.sql")
	c := string(path)

	db.MustExec(c)
}

func FillInDataBase(db *sqlx.DB) {

	scriptsPath := []string{"newSQLScripts/enums.sql", "newSQLScripts/podcast.sql", "newSQLScripts/_user_.sql",
		"newSQLScripts/album.sql", "newSQLScripts/updateUsers.sql"}

	for _, scriptPath := range scriptsPath {
		script, err := ioutil.ReadFile(scriptPath)

		if err != nil {
			panic(err)
		}

		stringScript := string(script)

		db.MustExec(stringScript)
	}
}

func GenerateMusicians(db *sqlx.DB) {
	var musicians []structs.RandomMusician

	f, err := os.Open("songs_normalize.csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read()

	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 400; i++ {
		var musician structs.RandomMusician
		row, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		musician = structs.RandomMusician{
			MusicianName: row[0],
		}

		musicians = append(musicians, musician)
	}

	GenerateMusician(db, musicians)
}

func GenerateTracks(db *sqlx.DB) {
	var tracks []structs.RandomTrack

	f, err := os.Open("songs_normalize.csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read()

	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 400; i++ {
		var track structs.RandomTrack
		row, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		explicit, _ := strconv.ParseBool(row[3])

		track = structs.RandomTrack{
			TrackName: row[1],
			Explicit:  explicit,
		}

		tracks = append(tracks, track)
	}

	GenerateTrack(db, tracks)
}

func GenerateData(db *sqlx.DB) {
	//GenerateMusicians(db)
	GenerateTracks(db)
	musicians, podcasts, tracks, users, albums := ParseDB(db)
	GeneratePlaylist(db, users)
	GeneratePodcastEP(db, podcasts)
	GenerateUserPlaylist(db, users)
	GenerateMusicianAlbum(db, musicians, albums)
	GenerateAlbumTrack(db, albums, tracks)
	GenerateUserMusician(db, users, musicians)
	GenerateUserAlbum(db, users, albums)
	GenerateUserTrack(db, users, tracks)
	GenerateUserPodcast(db, users, podcasts)
	GenerateUserPodcastEP(db, users)
	GenerateTrackPlaylist(db, tracks)
	GenerateSubscribes(db, users)
}
