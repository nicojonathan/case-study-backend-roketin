package parser

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
)

func ParseFormInsertUpdateMovie(r *http.Request) (request entity.InsertMovieRequest, err error) {
	err = formParser(r, &request)
	if err != nil {
		return entity.InsertMovieRequest{}, errors.New(err.Error())
	}

	return request, nil
}

func ParseFormSearchMovie(r *http.Request) (request entity.SearchMovieRequest, err error) {
	err = formParser(r, &request)
	if err != nil {
		return entity.SearchMovieRequest{}, errors.New(err.Error())
	}

	return request, nil
}

func ParseParamUpdateMovie(r *http.Request) (movieID int, err error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid URL path, missing movie ID")
	}

	idStr := parts[2]
	movieID, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid movie ID: %v", err)
	}
	return movieID, nil
}

func ParseVideoFile(r *http.Request, fieldName string) (multipart.File, *multipart.FileHeader, error) {
	maxUploadSize := 10 * 1024 * 1024
	err := r.ParseMultipartForm(int64(maxUploadSize))
	if err != nil {
		return nil, nil, err
	}

	file, fileHeader, err := r.FormFile(fieldName)
	if err != nil {
		return nil, nil, err
	}

	if fileHeader.Filename == "" {
		return nil, nil, errors.New("no file uploaded")
	}

	return file, fileHeader, nil
}
