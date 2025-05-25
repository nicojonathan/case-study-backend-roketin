package repository

import (
	"fmt"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/constant"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
)

func FindGenres(genreIDs []int) (genres []entity.Genre, err error) {
	db := connect()
	defer db.Close()

	if len(genreIDs) == 0 {
		return []entity.Genre{}, nil
	}

	// Build placeholders (?, ?, ?, ...)
	placeholders := make([]string, len(genreIDs))
	args := make([]interface{}, len(genreIDs))

	for i, id := range genreIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf("SELECT * FROM genres WHERE id IN (%s)", strings.Join(placeholders, ", "))

	rows, err := db.Query(query, args...)
	if err != nil {
		return []entity.Genre{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var genre entity.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return []entity.Genre{}, err
		}
		genres = append(genres, genre)
	}

	if len(genres) != len(genreIDs) {
		return []entity.Genre{}, fmt.Errorf(constant.NotFoundMessage)
	}

	if err = rows.Err(); err != nil {
		return []entity.Genre{}, err
	}

	return genres, nil
}
