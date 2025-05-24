package parser

import (
	"errors"
	"net/http"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
)

func ParseFormInsertMovie(r *http.Request) (request entity.InsertMovieRequest, err error) {
	err = formParser(r, &request)
	if err != nil {
		return entity.InsertMovieRequest{}, errors.New(err.Error())
	}

	return request, nil
}
