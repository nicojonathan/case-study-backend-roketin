package flow

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/constant"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/parser"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/repository"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
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

func SearchMovie(request entity.SearchMovieRequest) (movies []entity.MovieDetail, err error) {
	movies, err = repository.SearchMovie(request)
	if err != nil {
		return []entity.MovieDetail{}, err
	}

	return movies, nil
}

func UploadMovieToMongoDB(file multipart.File, fileHeader *multipart.FileHeader) (data entity.MovieMetadata, err error) {
	client, err := repository.ConnectMongo()
	if err != nil {
		return entity.MovieMetadata{}, fmt.Errorf("MongoDB connection failed")
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("movie")
	fsBucket, err := gridfs.NewBucket(db)
	if err != nil {
		return entity.MovieMetadata{}, fmt.Errorf("failed to create GridFS bucket")
	}

	// Upload file to GridFS
	uploadStream, err := fsBucket.OpenUploadStream(fileHeader.Filename)
	if err != nil {
		return entity.MovieMetadata{}, fmt.Errorf("failed to open upload stream")
	}
	defer uploadStream.Close()

	size, err := io.Copy(uploadStream, file)
	if err != nil {
		return entity.MovieMetadata{}, fmt.Errorf("failed to upload file")
	}

	data = entity.MovieMetadata{
		FileID:   uploadStream.FileID,
		Filename: fileHeader.Filename,
		Size:     size,
	}

	return data, nil
}
