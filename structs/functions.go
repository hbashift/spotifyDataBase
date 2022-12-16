package structs

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func RandomBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}

func RandomAmount() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(300-150) + 300
}

func RandomPlaylist(user User) (playlist PlaylistDTO) {
	playlist = PlaylistDTO{
		RandStringBytes(40),
		RandStringBytes(20),
		RandomBool(),
		RandStringBytes(20),
		RandomBool(),
		user.UserId,
		uuid.NewString(),
	}

	return
}

func RandomPodcastEp(podcast Podcast) (podcastEp PodcastEpisodeDTO) {
	podcastEp = PodcastEpisodeDTO{
		RandDate(),
		RandStringBytes(50),
		RandStringBytes(15),
		RandStringBytes(20),
		podcast.PodcastId,
		uuid.NewString(),
	}

	return
}
