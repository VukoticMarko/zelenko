package controller

import (
	"encoding/json"
	"net/http"
	"zelenko-backend/dto"
	"zelenko-backend/service"
)

type GreenObjectController interface {
	AddObject(response http.ResponseWriter, request *http.Request)
}

type greenObjectController struct{}

var (
	greenObjectService service.GreenObjectService
)

func NewGreenObjectController(service service.GreenObjectService) GreenObjectController {
	greenObjectService = service
	return &greenScoreController{}
}

func (*greenScoreController) AddObject(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var greenObject dto.IGreenObject

	err := json.NewDecoder(request.Body).Decode(&greenObject)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := greenObjectService.AddObject(greenObject)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
