package newGenerator

import (
	"db_seminar/structs"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"time"
)

func randMusician(db *sqlx.DB, musicians []structs.RandomMusician) *structs.MusicianDTO {
	var musician structs.MusicianDTO

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		musician = structs.RandMusician(musicians[rand.Intn(len(musicians))])

		rows, err := db.Queryx(existsMusicianQuery, musician.Nickname)

		if err != nil {
			log.Fatalln(err)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil
			}
		}

		err = rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}

	return &musician
}

func randTrack(db *sqlx.DB, tracks []structs.RandomTrack) *structs.TrackDTO {
	var track structs.TrackDTO

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())

		track = structs.RandTrack(tracks[rand.Intn(len(tracks))])

		rows, err := db.Queryx(existsTrackQuery, track.Title)

		if err != nil {
			log.Fatalln(err)
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil
			}
		}

		err = rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}

	return &track
}

func randomPlaylist(db *sqlx.DB, users []structs.User) (*structs.PlaylistDTO, error) {
	var user structs.User
	var playlist structs.PlaylistDTO

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user = users[rand.Intn(len(users))]

		playlist = structs.RandomPlaylist(user)

		rows, err := db.Queryx(existsPlaylistQuery, playlist.Description)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, err
			}
		}

		err = rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}

	return &playlist, nil
}

func randomPodcastEP(db *sqlx.DB, podcasts []structs.Podcast) (*structs.PodcastEpisodeDTO, error) {
	isExists := true
	podcastEP := structs.PodcastEpisodeDTO{}

	for isExists {
		rand.Seed(time.Now().UnixNano())
		podcast := podcasts[rand.Intn(len(podcasts))]

		podcastEP = structs.RandomPodcastEp(podcast)

		rows, err := db.Queryx(existsPodcastQuery, podcastEP.Description)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, err
			}
		}

		rows.Close()
	}

	return &podcastEP, nil
}

func randUserPlaylist(db *sqlx.DB, users []structs.User) (*structs.Playlist, *structs.User, error) {
	isExists := true
	var playlist structs.Playlist
	var user structs.User

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user = users[rand.Intn(len(users))]

		var playlists []structs.Playlist

		err := db.Select(&playlists, "SELECT playlist_id, user_id_fk FROM playlist")

		if err != nil {
			log.Fatalln(err)
		}

		playlist = playlists[rand.Intn(len(playlists))]

		rows, err0 := db.Queryx(existsUserPlaylistQuery, user.UserId, playlist.PlaylistId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &playlist, &user, nil
}

func randMusicianAlbum(db *sqlx.DB, musicians []structs.Musician, albums []structs.Album) (*structs.Musician,
	*structs.Album, error) {
	isExists := true
	var musician structs.Musician
	var album structs.Album

	for isExists {
		rand.Seed(time.Now().UnixNano())
		musician = musicians[rand.Intn(len(musicians))]
		album = albums[rand.Intn(len(albums))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT musician_id_fk, album_id_fk FROM musician_album WHERE musician_id_fk=$1 AND album_id_fk=$2)",
			musician.MusicianId, album.AlbumId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &musician, &album, nil
}

func randAlbumTrack(db *sqlx.DB, albums []structs.Album, tracks []structs.Track) (*structs.Album, *structs.Track,
	error) {
	var album structs.Album
	var track structs.Track

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		album = albums[rand.Intn(len(albums))]
		track = tracks[rand.Intn(len(tracks))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT album_id_fk, track_id_fk FROM album_track WHERE album_id_fk=$1 AND track_id_fk=$2)",
			album.AlbumId, track.TrackId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &album, &track, nil
}

func randUserMusician(db *sqlx.DB, users []structs.User, musicians []structs.Musician) (*structs.User,
	*structs.Musician, error) {
	var user structs.User
	var musician structs.Musician

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user = users[rand.Intn(len(users))]
		musician = musicians[rand.Intn(len(musicians))]

		rows, err := db.Queryx("SELECT EXISTS(SELECT user_id_fk, musician_id_fk FROM user_musician WHERE user_id_fk=$1 AND musician_id_fk=$2)",
			user.UserId, musician.MusicianId)

		if err != nil {
			return nil, nil, err
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &user, &musician, nil
}

func randUserAlbum(db *sqlx.DB, users []structs.User, albums []structs.Album) (*structs.User, *structs.Album, error) {
	var user structs.User
	var album structs.Album

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user = users[rand.Intn(len(users))]
		album = albums[rand.Intn(len(albums))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, album_id_fk FROM user_album WHERE user_id_fk=$1 AND album_id_fk=$2)",
			user.UserId, album.AlbumId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &user, &album, nil
}

func randUserTrack(db *sqlx.DB, users []structs.User, tracks []structs.Track) (*structs.User, *structs.Track, error) {
	var user structs.User
	var track structs.Track

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user = users[rand.Intn(len(users))]
		track = tracks[rand.Intn(len(users))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, track_id_fk FROM user_track WHERE user_id_fk=$1 AND track_id_fk=$2)",
			user.UserId, track.TrackId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &user, &track, nil
}

func randUserPodcast(db *sqlx.DB, users []structs.User, podcasts []structs.Podcast) (*structs.User, *structs.Podcast,
	error) {
	var user structs.User
	var podcast structs.Podcast

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user = users[rand.Intn(len(users))]
		podcast = podcasts[rand.Intn(len(podcasts))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, podcast_id_fk FROM user_podcast WHERE user_id_fk=$1 AND podcast_id_fk=$2)",
			user.UserId, podcast.PodcastId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &user, &podcast, nil
}

func randUserPodcastEp(db *sqlx.DB, users []structs.User) (*structs.User,
	*structs.PodcastEpisode, error) {
	var user structs.User
	var podcastEP structs.PodcastEpisode

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user = users[rand.Intn(len(users))]

		var podcastEPs []structs.PodcastEpisode

		err := db.Select(&podcastEPs, "SELECT podcast_ep_id, podcast_id_fk FROM podcast_episode")

		if err != nil {
			log.Fatalln(err)
		}

		podcastEP = podcastEPs[rand.Intn(len(podcastEPs))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_fk, podcast_ep_id_fk FROM user_podcast_ep WHERE user_id_fk=$1 AND podcast_ep_id_fk=$2)",
			user.UserId, podcastEP.PodcastEpId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &user, &podcastEP, nil
}

func randTrackPlaylist(db *sqlx.DB, tracks []structs.Track) (*structs.Track, *structs.PlaylistDTO, error) {
	var track structs.Track
	var playlist structs.PlaylistDTO

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		track = tracks[rand.Intn(len(tracks))]

		var playlists []structs.PlaylistDTO

		err := db.Select(&playlists, "SELECT * FROM playlist")

		if err != nil {
			log.Fatalln(err)
		}

		playlist = playlists[rand.Intn(len(playlists))]

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT track_id_fk, playlist_id_fk FROM track_playlist WHERE track_id_fk=$1 AND playlist_id_fk=$2)",
			track.TrackId, playlist.PlaylistId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &track, &playlist, nil
}

func randSubscribes(db *sqlx.DB, users []structs.User) (*structs.User, *structs.User, error) {
	var user1 structs.User
	var user2 structs.User

	isExists := true

	for isExists {
		rand.Seed(time.Now().UnixNano())
		user1 = users[rand.Intn(len(users))]
		user2 = users[rand.Intn(len(users))]

		CheckUsers(&user1, &user2, users)

		rows, err0 := db.Queryx("SELECT EXISTS(SELECT user_id_1, subscribes_user_id_2 FROM subscribes WHERE user_id_1=$1 AND subscribes_user_id_2=$2)",
			user1.UserId, user2.UserId)

		if err0 != nil {
			return nil, nil, err0
		}

		for rows.Next() {
			if err := rows.Scan(&isExists); err != nil {
				return nil, nil, err
			}
		}

		rows.Close()
	}

	return &user1, &user2, nil
}
