package model

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Для данных, если необходимо
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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

type AddedSong struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongsDetail struct {
	Text        string `json:"text"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
}

type SongPaginations struct {
	SongID int64    `json:"song_id"`
	Text   string   `json:"text"`
	Lines  []string `json:"lines"`
	Page   int      `json:"page"`
	Limit  int      `json:"limit"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type FilterParams struct {
	GroupName   string `json:"group_name"`
	Title       string `json:"title"`
	ReleaseDate string `json:"Release_date"`
	Text        string `json:"text"`
}

type LibraryRequest struct {
	Pagination Pagination   `json:"pagination"`
	Filters    FilterParams `json:"filters"`
}

type Library struct {
	Library []Input `json:"input"`
}
