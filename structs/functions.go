package structs

import (
	lorelai "github.com/UltiRequiem/lorelai/pkg"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	randomfiles "github.com/miihael/go-random-files"
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

	return rand.Intn(200-100) + 200
}

func RandMusician(musicianMeta RandomMusician) (musician MusicianDTO) {
	musician = MusicianDTO{
		musicianMeta.MusicianName,
		faker.Email(),
		randomfiles.RandomFilename(13, randomfiles.RunesEasy),
		uuid.NewString(),
	}

	return
}

func RandTrack(trackName RandomTrack) (track TrackDTO) {
	track = TrackDTO{
		trackName.TrackName,
		lorelai.FormattedLoremWords(3),
		lorelai.FormattedLoremWords(10),
		trackName.Explicit,
		randomfiles.RandomFilename(100, randomfiles.RunesHard),
		uuid.NewString(),
	}

	return
}

func RandomPlaylist(user User) (playlist PlaylistDTO) {

	playlist = PlaylistDTO{
		lorelai.FormattedLoremWords(10),
		randomfiles.RandomFilename(100, randomfiles.RunesHard),
		RandomBool(),
		lorelai.FormattedLoremWords(3),
		RandomBool(),
		user.UserId,
		uuid.NewString(),
	}

	return
}

func RandomPodcastEp(podcast Podcast) (podcastEp PodcastEpisodeDTO) {

	podcastEp = PodcastEpisodeDTO{
		RandDate(),
		lorelai.FormattedLoremWords(10),
		lorelai.FormattedLoremWords(3),
		randomfiles.RandomFilename(90, randomfiles.RunesHard),
		podcast.PodcastId,
		uuid.NewString(),
	}

	return
}
