package repository

import (
	"fmt"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/constant"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/parser"
)

func InsertMovie(request entity.Movie) (movieID int64, err error) {
	db := connect()
	defer db.Close()

	queryInsertMovie := "INSERT INTO movies (title, description, duration, video_file_id) VALUES (?, ?, ?, ?)"

	result, err := db.Exec(queryInsertMovie, request.Title, request.Description, request.Duration, request.VideoFileID)
	if err != nil {
		return 0, err
	}

	movieID, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return movieID, nil
}

func UpdateMovie(request entity.Movie) error {
	db := connect()
	defer db.Close()

	query := `
		UPDATE movies 
		SET title = ?, description = ?, duration = ?, video_file_id = ? 
		WHERE id = ?
	`

	_, err := db.Exec(query, request.Title, request.Description, request.Duration, request.VideoFileID, request.ID)
	if err != nil {
		return err
	}

	return nil
}

func InsertMovieArtist(movie entity.Movie, artists []entity.Artist) error {
	db := connect()
	defer db.Close()

	if len(artists) == 0 {
		return nil
	}

	if len(artists) == 1 {
		query := "INSERT INTO movie_artists (movie_id, artist_id) VALUES (?, ?)"
		_, err := db.Exec(query, movie.ID, artists[0].ID)
		return err
	}

	valueStrings := make([]string, 0, len(artists))
	valueArgs := make([]interface{}, 0, len(artists)*2)

	for _, artist := range artists {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, movie.ID, artist.ID)
	}

	query := fmt.Sprintf("INSERT INTO movie_artists (movie_id, artist_id) VALUES %s",
		strings.Join(valueStrings, ", "))

	_, err := db.Exec(query, valueArgs...)
	return err
}

func InsertMovieGenre(movie entity.Movie, genres []entity.Genre) error {
	db := connect()
	defer db.Close()

	if len(genres) == 0 {
		return nil
	}

	if len(genres) == 1 {
		query := "INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)"
		_, err := db.Exec(query, movie.ID, genres[0].ID)
		return err
	}

	valueStrings := make([]string, 0, len(genres))
	valueArgs := make([]interface{}, 0, len(genres)*2)

	for _, genre := range genres {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, movie.ID, genre.ID)
	}

	query := fmt.Sprintf("INSERT INTO movie_genres (movie_id, genre_id) VALUES %s",
		strings.Join(valueStrings, ", "))

	_, err := db.Exec(query, valueArgs...)
	return err
}

func DeleteMovieArtist(movieID int) error {
	db := connect()
	defer db.Close()

	query := "DELETE FROM movie_artists WHERE movie_id=?"

	_, err := db.Exec(query, movieID)
	return err
}

func DeleteMovieGenre(movieID int) error {
	db := connect()
	defer db.Close()

	query := "DELETE FROM movie_genres WHERE movie_id=?"

	_, err := db.Exec(query, movieID)
	return err
}

func GetAllMovies(limit int, page int) (movies []entity.MovieDetail, err error) {
	db := connect()
	defer db.Close()

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := `
		SELECT 
			m.id AS movie_id,
			m.title AS movie_title,
			m.description AS movie_description,
			m.duration AS movie_duration,
			m.video_file_id AS video_file_id,
			GROUP_CONCAT(DISTINCT a.name SEPARATOR ', ') AS artists,
			GROUP_CONCAT(DISTINCT g.name SEPARATOR ', ') AS genres
		FROM movies m
		LEFT JOIN movie_artists ma ON m.id = ma.movie_id
		LEFT JOIN artists a ON ma.artist_id = a.id
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		GROUP BY m.id
		ORDER BY m.id
		LIMIT ? OFFSET ?;
	`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return []entity.MovieDetail{}, err
	}

	defer rows.Close()

	moviesFound := false
	for rows.Next() {
		moviesFound = true
		var movie entity.MovieDetail
		rows.Scan(&movie.Movie.ID, &movie.Movie.Title, &movie.Movie.Description, &movie.Movie.Duration, &movie.Movie.VideoFileID, &movie.Artists, &movie.Genres)
		movies = append(movies, movie)
	}

	if !moviesFound {
		return []entity.MovieDetail{}, fmt.Errorf("movies " + constant.NotFoundMessage)
	}

	return movies, nil
}

func SearchMovie(request entity.SearchMovieRequest) (movies []entity.MovieDetail, err error) {
	db := connect()
	defer db.Close()

	baseQuery := `
		SELECT 
			m.id AS movie_id,
			m.title AS movie_title,
			m.description AS movie_description,
			m.duration AS movie_duration,
			m.video_file_id AS video_file_id,
			GROUP_CONCAT(DISTINCT a.name SEPARATOR ', ') AS artists,
			GROUP_CONCAT(DISTINCT g.name SEPARATOR ', ') AS genres
		FROM movies m
		LEFT JOIN movie_artists ma ON m.id = ma.movie_id
		LEFT JOIN artists a ON ma.artist_id = a.id
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
	`

	var conditions []string
	var args []interface{}

	if request.Title != "" {
		conditions = append(conditions, "LOWER(m.title) LIKE ?")
		args = append(args, "%"+strings.ToLower(request.Title)+"%")
	}
	if request.Description != "" {
		conditions = append(conditions, "LOWER(m.description) LIKE ?")
		args = append(args, "%"+strings.ToLower(request.Description)+"%")
	}
	if len(request.ArtistIDs) > 0 {
		arrArtistID, err := parser.ParseIDs(request.ArtistIDs)
		if err != nil {
			return []entity.MovieDetail{}, err
		}

		placeholders := strings.Repeat("?,", len(arrArtistID))
		placeholders = placeholders[:len(placeholders)-1] // remove comma
		conditions = append(conditions, fmt.Sprintf("ma.artist_id IN (%s)", placeholders))

		for _, id := range arrArtistID {
			args = append(args, id)
		}
	}
	if len(request.GenreIDs) > 0 {
		arrGenreID, err := parser.ParseIDs(request.GenreIDs)
		if err != nil {
			return []entity.MovieDetail{}, err
		}

		placeholders := strings.Repeat("?,", len(arrGenreID))
		placeholders = placeholders[:len(placeholders)-1]
		conditions = append(conditions, fmt.Sprintf("mg.genre_id IN (%s)", placeholders))

		for _, id := range arrGenreID {
			args = append(args, id)
		}
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	baseQuery += `
		GROUP BY m.id
		ORDER BY m.id;
	`
	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		return []entity.MovieDetail{}, err
	}
	defer rows.Close()

	moviesFound := false
	for rows.Next() {
		moviesFound = true
		var movie entity.MovieDetail
		rows.Scan(&movie.Movie.ID, &movie.Movie.Title, &movie.Movie.Description, &movie.Movie.Duration, &movie.Movie.VideoFileID, &movie.Artists, &movie.Genres)
		movies = append(movies, movie)
	}

	if !moviesFound {
		return []entity.MovieDetail{}, fmt.Errorf("movies " + constant.NotFoundMessage)
	}

	return movies, nil
}
