package response

import (
	"encoding/json"
	"net/http"

	"github.com/nicojonathan/case-study-backend-roketin/challenge2/entity"
)

func SendStandardSuccessResponse(w http.ResponseWriter, message string) {
	var response entity.StandardResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	var response entity.StandardResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponseWithData(w http.ResponseWriter, message string, data interface{}) {
	var response entity.ResponseWithData
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
