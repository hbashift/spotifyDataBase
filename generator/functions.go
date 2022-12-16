package generator

import (
	"db_seminar/structs"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	insertPlaylistQuery     = "INSERT INTO playlist VALUES ($1, $2, $3, $4, $5, $6, $7)"
	existsPlaylistQuery     = "SELECT EXISTS(SELECT description FROM playlist WHERE description = $1)"
	insertPodcastEpQuery    = "INSERT INTO podcast_episode VALUES ($1, $2, $3, $4, $5, $6)"
	existsUserPlaylistQuery = "SELECT EXISTS(SELECT user_id_fk, playlist_id_fk FROM user_playlist WHERE user_id_fk = $1 AND playlist_id_fk = $2)"
	existsPodcastQuery      = "SELECT EXISTS(SELECT description FROM podcast_episode WHERE description = $1)"
)

func GeneratePlaylist(db *sqlx.DB, users []structs.User) {
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		playlist, err := randomPlaylist(db, users)

		if err != nil {
			log.Fatalln(err)
		}

		row, err := db.Queryx(insertPlaylistQuery, playlist.Description, playlist.Cover,
			playlist.AccessModifier, playlist.Title, playlist.IsChart, playlist.UserIdFk, playlist.PlaylistId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GeneratePodcastEP(db *sqlx.DB, podcasts []structs.Podcast) {
	// Generating podcast_episode
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		podcastEP, err := randomPodcastEP(db, podcasts)

		row, err := db.Queryx(insertPodcastEpQuery, podcastEP.ReleaseDate,
			podcastEP.Description, podcastEP.Title, podcastEP.FileAddress, podcastEP.PodcastIdFk, podcastEP.PodcastEpId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateUserPlaylist(db *sqlx.DB, users []structs.User) {
	// Generating user_playlist
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		playlist, user, err0 := randUserPlaylist(db, users)

		if err0 != nil {
			log.Fatalln(err0)
		}

		row, err := db.Queryx("INSERT INTO user_playlist(user_id_fk, playlist_id_fk) VALUES ($1, $2)", user.UserId,
			playlist.PlaylistId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateMusicianAlbum(db *sqlx.DB, musicians []structs.Musician, albums []structs.Album) {
	// Generating musician_album
	amount := structs.RandomAmount()
	releaseDate := structs.RandDate()

	for i := 0; i < amount; i++ {
		musician, album, err := randMusicianAlbum(db, musicians, albums)

		if err != nil {
			log.Fatalln(err)
		}

		row, err := db.Queryx("INSERT INTO musician_album(release_date, musician_id_fk, album_id_fk) VALUES ($1, $2, $3)",
			releaseDate, musician.MusicianId, album.AlbumId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateAlbumTrack(db *sqlx.DB, albums []structs.Album, tracks []structs.Track) {
	// Generating album_track
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		album, track, err := randAlbumTrack(db, albums, tracks)

		if err != nil {
			log.Fatalln(err)
		}

		row, err := db.Queryx("INSERT INTO album_track(album_id_fk, track_id_fk) VALUES ($1, $2)",
			album.AlbumId, track.TrackId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateUserMusician(db *sqlx.DB, users []structs.User, musicians []structs.Musician) {
	// Generating user_musician
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		user, musician, err := randUserMusician(db, users, musicians)

		row, err := db.Queryx("INSERT INTO user_musician(user_id_fk, musician_id_fk) VALUES ($1, $2)",
			user.UserId, musician.MusicianId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateUserAlbum(db *sqlx.DB, users []structs.User, albums []structs.Album) {
	// Generating user_album
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		user, album, err := randUserAlbum(db, users, albums)

		if err != nil {
			log.Fatalln(err)
		}

		row, err := db.Queryx("INSERT INTO user_album(user_id_fk, album_id_fk) VALUES ($1, $2)", user.UserId,
			album.AlbumId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateUserTrack(db *sqlx.DB, users []structs.User, tracks []structs.Track) {
	// Generating user_track
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		likeDate := structs.RandDate()
		user, track, err := randUserTrack(db, users, tracks)

		row, err := db.Queryx("INSERT INTO user_track(like_date, user_id_fk, track_id_fk) VALUES ($1, $2, $3)",
			likeDate, user.UserId, track.TrackId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln()
		}
	}
}

func GenerateUserPodcast(db *sqlx.DB, users []structs.User, podcasts []structs.Podcast) {
	// Generating user_podcast
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		user, podcast, err := randUserPodcast(db, users, podcasts)

		if err != nil {
			log.Fatalln(err)
		}

		row, err := db.Queryx("INSERT INTO user_podcast(user_id_fk, podcast_id_fk) VALUES ($1, $2)",
			user.UserId, podcast.PodcastId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln()
		}
	}
}

func GenerateUserPodcastEP(db *sqlx.DB, users []structs.User) {
	// Generating user_podcast_ep
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		user, podcastEP, err := randUserPodcastEp(db, users)

		if err != nil {
			log.Fatalln(err)
		}

		row, err1 := db.Queryx("INSERT INTO user_podcast_ep(user_id_fk, podcast_ep_id_fk) VALUES ($1, $2)",
			user.UserId, podcastEP.PodcastEpId)

		if err1 != nil {
			log.Fatalln(err1)
		}

		err1 = row.Close()

		if err1 != nil {
			log.Fatalln(err1)
		}

	}
}

func GenerateTrackPlaylist(db *sqlx.DB, tracks []structs.Track) {
	// Generating track_playlist
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		track, playlist, err := randTrackPlaylist(db, tracks)

		if err != nil {
			log.Fatalln(err)
		}

		row, err1 := db.Queryx("INSERT INTO track_playlist(track_id_fk, playlist_id_fk) VALUES ($1, $2)", track.TrackId,
			playlist.PlaylistId)

		if err1 != nil {
			log.Fatalln(err1)
		}

		err1 = row.Close()

		if err1 != nil {
			log.Fatalln(err1)
		}
	}
}

func GenerateSubscribes(db *sqlx.DB, users []structs.User) {
	// Generating subscribes
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		user1, user2, err := randSubscribes(db, users)

		if err != nil {
			log.Fatalln(err)
		}

		row, err := db.Queryx("INSERT INTO subscribes(user_id_1, subscribes_user_id_2) VALUES ($1, $2)",
			user1.UserId, user2.UserId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}
