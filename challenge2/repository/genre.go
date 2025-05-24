package repository

import (
	"fmt"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
)

func FindGenres(gerneIDs []int) (genres []entity.Genre, err error) {
	db := connect()
	defer db.Close()

	// Handle empty input
	if len(gerneIDs) == 0 {
		return []entity.Genre{}, nil
	}

	// Build placeholders (?, ?, ?, ...)
	placeholders := make([]string, len(gerneIDs))
	args := make([]interface{}, len(gerneIDs))

	for i, id := range gerneIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	// Join the placeholders into the query
	query := fmt.Sprintf("SELECT * FROM gernes WHERE id IN (%s)", strings.Join(placeholders, ", "))

	rows, err := db.Query(query, args...)
	if err != nil {
		return []entity.Genre{}, err
	}
	defer rows.Close()

	// Fetch rows
	for rows.Next() {
		var genre entity.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}
