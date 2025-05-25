package flow

import (
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

	artists, err := repository.FindArtists(artistIDs)
	if err != nil {
		if err.Error() == constant.NotFoundMessage {
			return fmt.Errorf("Artist " + constant.NotFoundMessage)
		}
		return err
	}

	genres, err := repository.FindGenres(genreIDs)
	if err != nil {
		if err.Error() == constant.NotFoundMessage {
			return fmt.Errorf("Genre " + constant.NotFoundMessage)
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

func UpdateMovie(request entity.InsertMoviePayload) (err error) {
	artistIDs, err := parser.ParseIDs(request.ArtistIDs)
	if err != nil {
		return err
	}

	artists, err := repository.FindArtists(artistIDs)
	if err != nil {
		if err.Error() == constant.NotFoundMessage {
			return fmt.Errorf("Artist " + constant.NotFoundMessage)
		}
		return err
	}

	genreIDs, err := parser.ParseIDs(request.GenreIDs)
	if err != nil {
		return err
	}

	genres, err := repository.FindGenres(genreIDs)
	if err != nil {
		if err.Error() == constant.NotFoundMessage {
			return fmt.Errorf("Genre " + constant.NotFoundMessage)
		}
		return err
	}

	err = repository.UpdateMovie(request.Movie)
	if err != nil {
		return err
	}

	err = repository.DeleteMovieArtist(int(request.Movie.ID))
	if err != nil {
		return err
	}

	err = repository.DeleteMovieGenre(int(request.Movie.ID))
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

func GetAllMovies(request entity.GetAllMovieRequest) (movies []entity.MovieDetail, err error) {
	movies, err = repository.GetAllMovies(request.Limit, request.Page)
	if err != nil {
		return []entity.MovieDetail{}, err
	}

	return movies, nil
}
