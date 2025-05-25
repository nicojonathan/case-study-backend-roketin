package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/constant"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/flow"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/parser"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/response"
)

func InsertMovie(w http.ResponseWriter, r *http.Request) {
	request, err := parser.ParseFormInsertUpdateMovie(r)
	if err != nil {
		response.SendErrorResponse(w, 400, err.Error())
		return
	}

	requestJson, _ := json.Marshal(request)
	var payload entity.InsertMoviePayload
	_ = json.Unmarshal(requestJson, &payload)

	payload.Movie.ID = request.ID
	payload.Movie.Title = request.Title
	payload.Movie.Description = request.Description
	payload.Movie.Duration = request.Duration

	if request.ArtistIDs == "" || request.GenreIDs == "" || request.Title == "" || request.Description == "" {
		response.SendErrorResponse(w, 400, "Bad Request! All fields must be filled")

		return
	}

	err = flow.InsertMovie(payload)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFoundMessage) {
			response.SendErrorResponse(w, 404, err.Error())
			return
		}
		response.SendErrorResponse(w, 500, err.Error())
		return
	}

	response.SendStandardSuccessResponse(w, "Movies successfully inserted")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieID, err := parser.ParseParamUpdateMovie(r)
	if err != nil {
		response.SendErrorResponse(w, 400, err.Error())
		return
	}

	request, err := parser.ParseFormInsertUpdateMovie(r)
	if err != nil {
		response.SendErrorResponse(w, 400, err.Error())
		return
	}

	requestJson, _ := json.Marshal(request)
	var payload entity.InsertMoviePayload
	_ = json.Unmarshal(requestJson, &payload)

	payload.Movie.ID = int64(movieID)
	payload.Movie.Title = request.Title
	payload.Movie.Description = request.Description
	payload.Movie.Duration = request.Duration

	err = flow.UpdateMovie(payload)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFoundMessage) {
			response.SendErrorResponse(w, 404, err.Error())
			return
		}
		response.SendErrorResponse(w, 500, err.Error())
		return
	}

	response.SendStandardSuccessResponse(w, "Movies successfully updated")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	var request entity.GetAllMovieRequest

	params := parser.ParseQueryParams(r)
	request.Limit, _ = strconv.Atoi(params["limit"])
	request.Page, _ = strconv.Atoi(params["page"])

	movies, err := flow.GetAllMovies(request)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFoundMessage) {
			response.SendErrorResponse(w, 404, err.Error())
			return
		}
		response.SendErrorResponse(w, 500, err.Error())
		return
	}

	response.SendSuccessResponseWithData(w, "success", movies)
}

func SearchMovie(w http.ResponseWriter, r *http.Request) {
	request, err := parser.ParseFormSearchMovie(r)
	if err != nil {
		response.SendErrorResponse(w, 400, err.Error())
		return
	}

	movies, err := flow.SearchMovie(request)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFoundMessage) {
			response.SendErrorResponse(w, 404, err.Error())
			return
		}
		response.SendErrorResponse(w, 500, err.Error())
		return
	}

	response.SendSuccessResponseWithData(w, "success", movies)
}

func UploadMovieToMongoDB(w http.ResponseWriter, r *http.Request) {
	// limit the maximum file size that can be accepted by this endpoint
	r.Body = http.MaxBytesReader(w, r.Body, 2*1024*1024*1024)

	// Get file and its metadata
	file, fileHeader, err := parser.ParseVideoFile(r, "movie")
	if err != nil {
		response.SendErrorResponse(w, 400, err.Error())
		return
	}
	defer file.Close()

	data, err := flow.UploadMovieToMongoDB(file, fileHeader)
	if err != nil {
		response.SendErrorResponse(w, 500, err.Error())
		return
	}

	response.SendSuccessResponseWithData(w, "Movie has been uploaded successfully", data)
}
