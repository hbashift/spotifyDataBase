package generator

import (
	"db_seminar/structs"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"time"
)

func GeneratePlaylist(db *sqlx.DB, users []structs.User) {
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		user := users[rand.Intn(len(users))]

		playlist := structs.RandomPlaylist(user)

		row, err := db.Queryx("INSERT INTO playlist VALUES ($1, $2, $3, $4, $5, $6, $7)", playlist.Description,
			playlist.Cover, playlist.AccessModifier, playlist.Title,
			playlist.IsChart, playlist.UserIdFk, playlist.PlaylistId)

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
		rand.Seed(time.Now().UnixNano())
		podcast := podcasts[rand.Intn(len(podcasts))]

		podcastEP := structs.RandomPodcastEp(podcast)

		row, err := db.Queryx("INSERT INTO podcast_episode VALUES ($1, $2, $3, $4, $5, $6)", podcastEP.ReleaseDate,
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
		var isExists bool

		rand.Seed(time.Now().UnixNano())
		user := users[rand.Intn(len(users))]

		var playlists []structs.PlaylistDTO

		err := db.Select(&playlists, "SELECT * FROM playlist")

		if err != nil {
			log.Fatalln(err)
		}

		playlist := playlists[rand.Intn(len(playlists))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, playlist_id_fk FROM user_playlist WHERE user_id_fk = $1 AND playlist_id_fk = $2)",
			user.UserId, playlist.PlaylistId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			user = users[rand.Intn(len(users))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT user_id_fk, playlist_id_fk FROM user_playlist WHERE user_id_fk = $1 AND playlist_id_fk = $2)",
				user.UserId, playlist.PlaylistId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err1 := db.Queryx("INSERT INTO user_playlist(user_id_fk, playlist_id_fk) VALUES ($1, $2)", user.UserId,
			playlist.PlaylistId)

		if err1 != nil {
			log.Fatalln(err1)
		}

		err1 = row.Close()
		rows.Close()

		if err1 != nil {
			log.Fatalln(err1)
		}
	}
}

func GenerateMusicianAlbum(db *sqlx.DB, musicians []structs.Musician, albums []structs.Album) {
	// Generating musician_album
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		musician := musicians[rand.Intn(len(musicians))]
		album := albums[rand.Intn(len(albums))]
		releaseDate := structs.RandDate()

		var isExists bool

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT musician_id_fk, album_id_fk FROM musician_album WHERE musician_id_fk=$1 AND album_id_fk=$2)",
			musician.MusicianId, album.AlbumId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			musician = musicians[rand.Intn(len(musicians))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT musician_id_fk, album_id_fk FROM musician_album WHERE musician_id_fk=$1 AND album_id_fk=$2)",
				musician.MusicianId, album.AlbumId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err := db.Queryx("INSERT INTO musician_album(release_date, musician_id_fk, album_id_fk) VALUES ($1, $2, $3)",
			releaseDate, musician.MusicianId, album.AlbumId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()
		rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateAlbumTrack(db *sqlx.DB, albums []structs.Album, tracks []structs.Track) {
	// Generating album_track
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		album := albums[rand.Intn(len(albums))]
		track := tracks[rand.Intn(len(tracks))]

		var isExists bool

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT album_id_fk, track_id_fk FROM album_track WHERE album_id_fk=$1 AND track_id_fk=$2)",
			album.AlbumId, track.TrackId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			album = albums[rand.Intn(len(albums))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT album_id_fk, track_id_fk FROM album_track WHERE album_id_fk=$1 AND track_id_fk=$2)",
				album.AlbumId, track.TrackId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err := db.Queryx("INSERT INTO album_track(album_id_fk, track_id_fk) VALUES ($1, $2)",
			album.AlbumId, track.TrackId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()
		rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateUserMusician(db *sqlx.DB, users []structs.User, musicians []structs.Musician) {
	// Generating user_musician
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		user := users[rand.Intn(len(users))]
		musician := musicians[rand.Intn(len(musicians))]

		var isExists bool

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, musician_id_fk FROM user_musician WHERE user_id_fk=$1 AND musician_id_fk=$2)",
			user.UserId, musician.MusicianId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			user = users[rand.Intn(len(users))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT user_id_fk, musician_id_fk FROM user_musician WHERE user_id_fk=$1 AND musician_id_fk=$2)",
				user.UserId, musician.MusicianId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err := db.Queryx("INSERT INTO user_musician(user_id_fk, musician_id_fk) VALUES ($1, $2)",
			user.UserId, musician.MusicianId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()
		rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateUserAlbum(db *sqlx.DB, users []structs.User, albums []structs.Album) {
	// Generating user_album
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		user := users[rand.Intn(len(users))]
		album := albums[rand.Intn(len(albums))]

		var isExists bool

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, album_id_fk FROM user_album WHERE user_id_fk=$1 AND album_id_fk=$2)",
			user.UserId, album.AlbumId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			user = users[rand.Intn(len(users))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT user_id_fk, album_id_fk FROM user_album WHERE user_id_fk=$1 AND album_id_fk=$2)",
				user.UserId, album.AlbumId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err := db.Queryx("INSERT INTO user_album(user_id_fk, album_id_fk) VALUES ($1, $2)", user.UserId,
			album.AlbumId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()
		rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GenerateUserTrack(db *sqlx.DB, users []structs.User, tracks []structs.Track) {
	// Generating user_track
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		user := users[rand.Intn(len(users))]
		track := tracks[rand.Intn(len(users))]
		likeDate := structs.RandDate()

		var isExists bool

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, track_id_fk FROM user_track WHERE user_id_fk=$1 AND track_id_fk=$2)",
			user.UserId, track.TrackId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			user = users[rand.Intn(len(users))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT user_id_fk, track_id_fk FROM user_track WHERE user_id_fk=$1 AND track_id_fk=$2)",
				user.UserId, track.TrackId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err := db.Queryx("INSERT INTO user_track(like_date, user_id_fk, track_id_fk) VALUES ($1, $2, $3)",
			likeDate, user.UserId, track.TrackId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()
		rows.Close()

		if err != nil {
			log.Fatalln()
		}
	}
}

func GenerateUserPodcast(db *sqlx.DB, users []structs.User, podcasts []structs.Podcast) {
	// Generating user_podcast
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		user := users[rand.Intn(len(users))]
		podcast := podcasts[rand.Intn(len(podcasts))]

		var isExists bool

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, podcast_id_fk FROM user_podcast WHERE user_id_fk=$1 AND podcast_id_fk=$2)",
			user.UserId, podcast.PodcastId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			user = users[rand.Intn(len(users))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT user_id_fk, podcast_id_fk FROM user_podcast WHERE user_id_fk=$1 AND podcast_id_fk=$2)",
				user.UserId, podcast.PodcastId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err := db.Queryx("INSERT INTO user_podcast(user_id_fk, podcast_id_fk) VALUES ($1, $2)",
			user.UserId, podcast.PodcastId)

		if err != nil {
			log.Fatalln(err)
		}

		err = row.Close()
		rows.Close()

		if err != nil {
			log.Fatalln()
		}
	}
}

func GenerateUserPodcastEP(db *sqlx.DB, users []structs.User) {
	// Generating user_podcast_ep
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		user := users[rand.Intn(len(users))]

		var isExists bool
		var podcastEPs []structs.PodcastEpisode

		err := db.Select(&podcastEPs, "SELECT podcast_ep_id, podcast_id_fk FROM podcast_episode")

		if err != nil {
			log.Fatalln(err)
		}

		podcastEP := podcastEPs[rand.Intn(len(podcastEPs))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, podcast_ep_id_fk FROM user_podcast_ep WHERE user_id_fk=$1 AND podcast_ep_id_fk=$2)",
			user.UserId, podcastEP.PodcastEpId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			user = users[rand.Intn(len(users))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT user_id_fk, podcast_ep_id_fk FROM user_podcast_ep WHERE user_id_fk=$1 AND podcast_ep_id_fk=$2)",
				user.UserId, podcastEP.PodcastEpId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err1 := db.Queryx("INSERT INTO user_podcast_ep(user_id_fk, podcast_ep_id_fk) VALUES ($1, $2)",
			user.UserId, podcastEP.PodcastEpId)

		if err1 != nil {
			log.Fatalln(err1)
		}

		err1 = row.Close()
		rows.Close()

		if err1 != nil {
			log.Fatalln(err1)
		}

	}
}

func GenerateTrackPlaylist(db *sqlx.DB, tracks []structs.Track) {
	// Generating track_playlist
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		track := tracks[rand.Intn(len(tracks))]

		var isExists bool
		var playlists []structs.PlaylistDTO

		err := db.Select(&playlists, "SELECT * FROM playlist")

		if err != nil {
			log.Fatalln(err)
		}

		playlist := playlists[rand.Intn(len(playlists))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT track_id_fk, playlist_id_fk FROM track_playlist WHERE track_id_fk=$1 AND playlist_id_fk=$2)",
			track.TrackId, playlist.PlaylistId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			track = tracks[rand.Intn(len(tracks))]

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT track_id_fk, playlist_id_fk FROM track_playlist WHERE track_id_fk=$1 AND playlist_id_fk=$2)",
				track.TrackId, playlist.PlaylistId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}

			rows.Close()
		}

		row, err1 := db.Queryx("INSERT INTO track_playlist(track_id_fk, playlist_id_fk) VALUES ($1, $2)", track.TrackId,
			playlist.PlaylistId)

		if err1 != nil {
			log.Fatalln(err1)
		}

		err1 = row.Close()
		rows.Close()

		if err1 != nil {
			log.Fatalln(err1)
		}
	}
}

func GenerateSubscribes(db *sqlx.DB, users []structs.User) {
	// Generating subscribes
	amount := structs.RandomAmount()

	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		user1 := users[rand.Intn(len(users))]
		user2 := users[rand.Intn(len(users))]
		var isExists bool

		CheckUsers(&user1, &user2, users)

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_1, subscribes_user_id_2 FROM subscribes WHERE user_id_1=$1 AND subscribes_user_id_2=$2)",
			user1.UserId, user2.UserId)

		if err0 != nil {
			log.Fatalln(err0)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				log.Fatalln(err)
			}
		}

		for isExists == true {
			rand.Seed(time.Now().UnixNano())
			user1 = users[rand.Intn(len(users))]

			CheckUsers(&user1, &user2, users)

			rows, err0 = db.Queryx("SELECT EXISTS(SELECT user_id_1, subscribes_user_id_2 FROM subscribes WHERE user_id_1=$1 AND subscribes_user_id_2=$2)",
				user1.UserId, user2.UserId)

			if err0 != nil {
				log.Fatalln(err0)
			}

			for rows.Next() {
				if err := rows.Scan(&isExists); err != nil {
					log.Fatalln(err)
				}
			}
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
