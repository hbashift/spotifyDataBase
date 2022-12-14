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
	enums, _ := ioutil.ReadFile("sqlScripts/enums.sql")
	sqlScript := string(enums)

	db.MustExec(sqlScript)

	musicians, _ := ioutil.ReadFile("sqlScripts/musician.sql")
	sqlScript = string(musicians)

	db.MustExec(sqlScript)

	podcast, _ := ioutil.ReadFile("sqlScripts/podcast.sql")
	sqlScript = string(podcast)

	db.MustExec(sqlScript)

	track, _ := ioutil.ReadFile("sqlScripts/track.sql")
	sqlScript = string(track)

	db.MustExec(sqlScript)

	podcastEp, _ := ioutil.ReadFile("sqlScripts/podcast_episode.sql")
	sqlScript = string(podcastEp)

	db.MustExec(sqlScript)

	user, _ := ioutil.ReadFile("sqlScripts/_user_.sql")
	sqlScript = string(user)

	db.MustExec(sqlScript)

	playlist, _ := ioutil.ReadFile("sqlScripts/playlist.sql")
	sqlScript = string(playlist)

	db.MustExec(sqlScript)

	album, _ := ioutil.ReadFile("sqlScripts/album.sql")
	sqlScript = string(album)

	db.MustExec(sqlScript)

	userPlaylist, _ := ioutil.ReadFile("sqlScripts/user_playlist.sql")
	sqlScript = string(userPlaylist)

	db.MustExec(sqlScript)

	musicianAlbum, _ := ioutil.ReadFile("sqlScripts/musician_album.sql")
	sqlScript = string(musicianAlbum)

	db.MustExec(sqlScript)

	albumTrack, _ := ioutil.ReadFile("sqlScripts/album_track.sql")
	sqlScript = string(albumTrack)

	db.MustExec(sqlScript)

	userMusician, _ := ioutil.ReadFile("sqlScripts/user_musician.sql")
	sqlScript = string(userMusician)

	db.MustExec(sqlScript)

	userAlbum, _ := ioutil.ReadFile("sqlScripts/user_album.sql")
	sqlScript = string(userAlbum)

	db.MustExec(sqlScript)

	userTrack, _ := ioutil.ReadFile("sqlScripts/user_track.sql")
	sqlScript = string(userTrack)

	db.MustExec(sqlScript)

	userPodcast, _ := ioutil.ReadFile("sqlScripts/user_podcast.sql")
	sqlScript = string(userPodcast)

	db.MustExec(sqlScript)

	userPodcastEp, _ := ioutil.ReadFile("sqlScripts/user_podcast_ep.sql")
	sqlScript = string(userPodcastEp)

	db.MustExec(sqlScript)

	trackPlaylist, _ := ioutil.ReadFile("sqlScripts/track_playlist.sql")
	sqlScript = string(trackPlaylist)

	db.MustExec(sqlScript)

	subscribes, _ := ioutil.ReadFile("sqlScripts/subscribes.sql")
	sqlScript = string(subscribes)

	db.MustExec(sqlScript)
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
