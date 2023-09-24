package controller

import (
	"encoding/json"
	"net/http"
	"zelenko-backend/model"
	"zelenko-backend/service"
)

type GreenScoreController interface {
	AddOne(response http.ResponseWriter, request *http.Request)
	SubOne(response http.ResponseWriter, request *http.Request)
}

type greenScoreController struct{}

var (
	greenScoreService service.GreenScoreService
)

func NewGreenScoreController(service service.GreenScoreService) GreenScoreController {
	greenScoreService = service
	return &greenScoreController{}
}

func (*greenScoreController) AddOne(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var greenObject model.GreenObject
	greenObject = greenScoreService.AddOne()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(greenObject)
}

func (*greenScoreController) SubOne(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var greenObject model.GreenObject
	greenObject = greenScoreService.SubOne()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(greenObject)
}
