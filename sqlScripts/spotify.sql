CREATE TABLE musician
(
  nickname varchar(128) NOT NULL,
  email varchar(128) NOT NULL UNIQUE CHECK (email LIKE '%_@_%_.__%'),
  password varchar(128) NOT NULL,
  musician_id uuid PRIMARY KEY
);

CREATE TABLE subscription_type
(
  subscription_title varchar(128) NOT NULL,
  sub_id int PRIMARY KEY
);

CREATE TABLE podcast
(
  author varchar(128) NOT NULL,
  description varchar(512) DEFAULT NULL,
  podcast_id uuid PRIMARY KEY
);

CREATE TABLE track
(
  title varchar(128) NOT NULL,
  lyrics_author varchar(100) DEFAULT NULL,
  lyrics text DEFAULT NULL,
  explicit boolean NOT NULL DEFAULT false,
  file_address varchar(512) NOT NULL UNIQUE,
  track_id uuid PRIMARY KEY
);

CREATE TABLE podcast_episode
(
  release_date date NOT NULL,
  description varchar(512) DEFAULT NULL,
  title varchar(128) NOT NULL,
  file_address varchar(512) NOT NULL,
  podcast_id_fk uuid NOT NULL REFERENCES podcast(podcast_id),
  podcast_ep_id uuid PRIMARY KEY
);

CREATE TABLE album_type
(
  type_name varchar(128) NOT NULL,
  type_id int PRIMARY KEY
);

CREATE TABLE "user"
(
  bank_card_number int DEFAULT NULL,
  subscription_expiration_date date DEFAULT NULL,
  email varchar(128) NOT NULL UNIQUE CHECK (email LIKE '%_@_%_.__%'),
  password varchar(128) NOT NULL,
  sub_id_fk int NOT NULL REFERENCES subscription_type(sub_id),
  user_id uuid PRIMARY KEY
);

CREATE TABLE playlist
(
  description varchar(256) DEFAULT NULL,
  cover text DEFAULT NULL,
  access_modifier boolean NOT NULL,
  title varchar(128) NOT NULL DEFAULT 'NewPlaylist',
  is_chart boolean DEFAULT false,
  user_id_fk uuid NOT NULL REFERENCES "user"(user_id),
  playlist_id uuid PRIMARY KEY
);

CREATE TABLE album
(
  title varchar(512) NOT NULL,
  cover text NOT NULL,
  label varchar(128) NOT NULL,
  type_id_fk int NOT NULL REFERENCES album_type(type_id),
  album_id uuid PRIMARY KEY
);

CREATE TABLE user_playlist
(
  user_id_fk uuid NOT NULL REFERENCES "user"(user_id),
  playlist_id_fk uuid NOT NULL REFERENCES playlist(playlist_id),
  UNIQUE (user_id_fk, playlist_id_fk),
  user_playlist_id serial PRIMARY KEY
);

CREATE TABLE musician_album
(
  release_date date NOT NULL,
  musician_id_fk uuid NOT NULL REFERENCES musician(musician_id),
  album_id_fk uuid NOT NULL REFERENCES album(album_id),
  UNIQUE (musician_id_fk, album_id_fk),
  musician_album_id serial PRIMARY KEY
);

CREATE TABLE album_track
(
  album_id_fk uuid NOT NULL REFERENCES album(album_id),
  track_id_fk uuid NOT NULL REFERENCES track(track_id),
  UNIQUE (album_id_fk, track_id_fk),
  album_track_id serial PRIMARY KEY
);

CREATE TABLE user_musician
(
  user_id_fk uuid NOT NULL REFERENCES "user"(user_id),
  musician_id_fk uuid NOT NULL REFERENCES musician(musician_id),
  UNIQUE (user_id_fk, musician_id_fk),
  user_musician_id serial PRIMARY KEY
);

CREATE TABLE user_album
(
  user_id_fk uuid NOT NULL REFERENCES "user"(user_id),
  album_id_fk uuid NOT NULL REFERENCES album(album_id),
  UNIQUE (user_id_fk, album_id_fk),
  user_album_id serial PRIMARY KEY
);

CREATE TABLE user_track
(
  like_date date NOT NULL,
  user_id_fk uuid NOT NULL REFERENCES "user"(user_id),
  track_id_fk uuid NOT NULL REFERENCES track(track_id),
  UNIQUE (user_id_fk, track_id_fk),
  user_track_id serial PRIMARY KEY
);

CREATE TABLE user_podcast
(
  user_id_fk uuid NOT NULL REFERENCES "user"(user_id),
  podcast_id_fk uuid NOT NULL REFERENCES podcast(podcast_id),
  UNIQUE (user_id_fk, podcast_id_fk),
  user_podcast_id serial PRIMARY KEY
);

CREATE TABLE user_podcast_ep
(
  user_id_fk uuid NOT NULL REFERENCES "user"(user_id),
  podcast_ep_id_fk uuid  NOT NULL REFERENCES podcast_episode(podcast_ep_id),
  UNIQUE (user_id_fk, podcast_ep_id_fk),
  user_podcast_ep_id serial PRIMARY KEY
);

CREATE TABLE track_playlist
(
  track_id_fk uuid NOT NULL REFERENCES track(track_id),
  playlist_id_fk uuid NOT NULL REFERENCES playlist(playlist_id),
  track_playlist_id serial PRIMARY KEY
);

CREATE TABLE subscribes
(
  user_id_1 uuid NOT NULL REFERENCES "user"(user_id),
  subscribes_user_id_2 uuid NOT NULL REFERENCES "user"(user_id),
  UNIQUE (user_id_1, subscribes_user_id_2),
  subscribes_id serial PRIMARY KEY
);
