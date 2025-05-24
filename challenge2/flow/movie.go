package flow

import (
	"database/sql"
	"fmt"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/constant"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/parser"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/repository"
)

func InsertMovie(request entity.InsertMoviePayload) error {
	artistIDs, err := parser.ParseIDs(request.ArtistIDs)
	if err != nil {
		return err
	}

	genreIDs, err := parser.ParseIDs(request.GenreIDs)
	if err != nil {
		return err
	}

	// get artist based on id
	artists, err := repository.FindArtists(artistIDs)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Artist" + constant.NotFoundMessage)
		}
		return err
	}

	genres, err := repository.FindGenres(genreIDs)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Genre" + constant.NotFoundMessage)
		}
		return err
	}

	request.Movie.ID, err = repository.InsertMovie(request.Movie)
	if err != nil {
		return err
	}

	err = repository.InsertMovieArtist(request.Movie, artists)
	if err != nil {
		return err
	}

	err = repository.InsertMovieGenre(request.Movie, genres)
	if err != nil {
		return err
	}

	return nil
}
