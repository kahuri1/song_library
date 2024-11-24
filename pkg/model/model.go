package model

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

type Group struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Song struct {
	SongID      int64  `json:"song_id"`
	Title       string `json:"title"`
	GroupID     int64  `json:"group_id"`
	ReleaseDate string `json:"release_date"`
	Lyrics      string `json:"lyrics"`
	Link        string `json:"link"`
}

type Input struct {
	Group Group `json:"group"`
	Song  Song  `json:"song"`
}
