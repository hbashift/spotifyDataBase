package structs

import (
	"time"
)

type Musician struct {
	MusicianId string `db:"musician_id"`
}

type SubscriptionType struct {
	SubId int `db:"sub_id"`
}

type Podcast struct {
	PodcastId string `db:"podcast_id"`
}

type Track struct {
	TrackId string `db:"track_id"`
}

type RandomTrack struct {
	TrackName string `csv:"song"`
	Explicit  bool   `db:"explicit"`
}

type RandomMusician struct {
	MusicianName string `csv:"artist"`
}

type User struct {
	UserId string `db:"user_id"`
}

type PodcastEpisode struct {
	PodcastIdFk string `db:"podcast_id_fk"`
	PodcastEpId string `db:"podcast_ep_id"`
}

type Playlist struct {
	UserIdFk   string `db:"user_id_fk"`
	PlaylistId string `db:"playlist_id"`
}

type Album struct {
	AlbumId string `db:"album_id"`
}

type PlaylistDTO struct {
	Description    string `db:"description"`
	Cover          string `db:"cover"`
	AccessModifier bool   `db:"access_modifier"`
	Title          string `db:"title"`
	IsChart        bool   `db:"is_chart"`
	UserIdFk       string `db:"user_id_fk"`
	PlaylistId     string `db:"playlist_id"`
}

type PodcastEpisodeDTO struct {
	ReleaseDate time.Time
	Description string
	Title       string
	FileAddress string
	PodcastIdFk string
	PodcastEpId string
}

type TrackDTO struct {
	Title        string `db:"title"`
	LyricsAuthor string `db:"lyrics_author"`
	Lyrics       string `db:"lyrics"`
	Explicit     bool   `db:"explicit"`
	FileAddress  string `db:"file_address"`
	TrackId      string `db:"track_id"`
}

type MusicianDTO struct {
	Nickname   string `db:"nickname"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	MusicianId string `db:"musician_id"`
}
