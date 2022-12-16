package generator

import (
	"db_seminar/structs"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"math/rand"
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
	path, _ := ioutil.ReadFile("sqlScripts/spotify.sql")
	c := string(path)

	db.MustExec(c)
}

func FillInDataBase(db *sqlx.DB) {

	scriptsPath := []string{"sqlScripts/enums.sql", "sqlScripts/musician.sql", "sqlScripts/podcast.sql",
		"sqlScripts/track.sql", "sqlScripts/_user_.sql", "sqlScripts/album.sql", "sqlScripts/updateUsers.sql"}

	for _, scriptPath := range scriptsPath {
		script, err := ioutil.ReadFile(scriptPath)

		if err != nil {
			panic(err)
		}

		stringScript := string(script)

		db.MustExec(stringScript)
	}
}

func GenerateData(db *sqlx.DB) {

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
