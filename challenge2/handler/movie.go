package handler

import (
	"net/http"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/constant"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/flow"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/parser"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/response"
)

func InsertMovie(w http.ResponseWriter, r *http.Request) {
	request, err := parser.ParseFormInsertMovie(r)
	if err != nil {
		response.SendErrorResponse(w, 400, err.Error())
		return
	}

	if request.ArtistIDs == "" || request.GenreIDs == "" {
		response.SendErrorResponse(w, 400, "Bad Request! Movie Category and Movie Genre can't be empty")
	}

	err = flow.InsertMovie(request)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFoundMessage) {
			response.SendErrorResponse(w, 404, err.Error())
		}
		response.SendErrorResponse(w, 500, err.Error())
		return
	}

	response.SendPostSuccessResponse(w, "Movies successfully inserted")
}
