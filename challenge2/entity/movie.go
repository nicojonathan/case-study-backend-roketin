package entity

type InsertMovieRequest struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	ArtistIDs   string `json:"artist_ids"`
	GenreIDs    string `json:"genre_ids"`
}

type GetAllMovieRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type SearchMovieRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ArtistIDs   string `json:"artist_ids"`
	GenreIDs    string `json:"genre_ids"`
}

type InsertMoviePayload struct {
	Movie     Movie  `json:"movie"`
	ArtistIDs string `json:"artist_ids"`
	GenreIDs  string `json:"genre_ids"`
}

type Movie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
}

type MovieDetail struct {
	Movie   Movie  `json:"movie"`
	Artists string `json:"artists"`
	Genres  string `json:"genres"`
}

type MovieMetadata struct {
	FileID   interface{} `json:"file_id"`
	Filename string      `json:"filename"`
	Size     int64       `json:"size"`
}
