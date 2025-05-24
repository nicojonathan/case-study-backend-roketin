package entity

type InsertMovieRequest struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	ArtistIDs   string `json:"artist_ids"`
	GenreIDs    string `json:"genres_ids"`
}

type InsertMoviePayload struct {
	Movie     Movie  `json:"movie"`
	ArtistIDs string `json:"artist_ids"`
	GenreIDs  string `json:"genres_ids"`
}

type Movie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
}

type StandardResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
