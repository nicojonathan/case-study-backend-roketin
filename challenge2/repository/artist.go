package repository

import (
	"fmt"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
)

func FindArtists(artistIDs []int) (artists []entity.Artist, err error) {
	db := connect()
	defer db.Close()

	// Handle empty input
	if len(artistIDs) == 0 {
		return []entity.Artist{}, nil
	}

	// Build placeholders (?, ?, ?, ...)
	placeholders := make([]string, len(artistIDs))
	args := make([]interface{}, len(artistIDs))

	for i, id := range artistIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	// Join the placeholders into the query
	query := fmt.Sprintf("SELECT * FROM artists WHERE id IN (%s)", strings.Join(placeholders, ", "))

	rows, err := db.Query(query, args...)
	if err != nil {
		return []entity.Artist{}, err
	}
	defer rows.Close()

	// Fetch rows
	for rows.Next() {
		var artist entity.Artist
		if err := rows.Scan(&artist.ID, &artist.Name); err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}
