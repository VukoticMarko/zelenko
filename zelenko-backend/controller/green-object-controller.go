package controller

import (
	"encoding/json"
	"net/http"
	"zelenko-backend/dto"
	"zelenko-backend/service"
)

type GreenObjectController interface {
	AddObject(response http.ResponseWriter, request *http.Request)
	GetAllObjects(response http.ResponseWriter, request *http.Request)
	UpdateObject(response http.ResponseWriter, request *http.Request)
	DeleteObject(response http.ResponseWriter, request *http.Request)
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

func (*greenScoreController) GetAllObjects(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	result := greenObjectService.FindAll()

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*greenScoreController) UpdateObject(response http.ResponseWriter, request *http.Request) {

	var greenObject dto.IGreenObject

	err := json.NewDecoder(request.Body).Decode(&greenObject)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := greenObjectService.UpdateObject(greenObject)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*greenScoreController) DeleteObject(response http.ResponseWriter, request *http.Request) {

	var greenObject dto.IGreenObject

	err := json.NewDecoder(request.Body).Decode(&greenObject)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	greenObjectService.DeleteObject(greenObject)

	response.WriteHeader(http.StatusOK)
}
