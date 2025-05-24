package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/constant"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/flow"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/parser"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/response"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/util"
)

func InsertMovie(w http.ResponseWriter, r *http.Request) {
	request, err := parser.ParseFormInsertMovie(r)
	if err != nil {
		response.SendErrorResponse(w, 400, err.Error())
		return
	}

	util.PrintJSON(request)

	if request.ArtistIDs == "" || request.GenreIDs == "" || request.Title == "" || request.Description == "" {
		response.SendErrorResponse(w, 400, "Bad Request! All fields must be filled")

		return
	}

	requestJson, _ := json.Marshal(request)
	var payload entity.InsertMoviePayload
	_ = json.Unmarshal(requestJson, &payload)

	payload.Movie.ID = request.ID
	payload.Movie.Title = request.Title
	payload.Movie.Description = request.Description
	payload.Movie.Duration = request.Duration

	err = flow.InsertMovie(payload)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFoundMessage) {
			response.SendErrorResponse(w, 404, err.Error())
		}
		response.SendErrorResponse(w, 500, err.Error())
		return
	}

	response.SendPostSuccessResponse(w, "Movies successfully inserted")
}
