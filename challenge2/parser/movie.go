package parser

import (
	"errors"
	"fmt"
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
