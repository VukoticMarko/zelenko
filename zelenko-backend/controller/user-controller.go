package controller

import (
	"encoding/json"
	"net/http"
	"zelenko-backend/model"
	"zelenko-backend/service"
)

type UserController interface {
	AddUser(response http.ResponseWriter, request *http.Request)
	//UpdateUser(response http.ResponseWriter, request *http.Request)
	GetUser(response http.ResponseWriter, request *http.Request)
	//GetUsers(response http.ResponseWriter, request *http.Request)
	//DeleteUser(response http.ResponseWriter, request *http.Request)
}

type userController struct{}

var (
	userService service.UserService
)

func NewUserController(service service.UserService) UserController {
	userService = service
	return &userController{}
}

func (*userController) AddUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var req model.User

	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err1 := userService.AddUser(req)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*userController) GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var greenObject model.GreenObject
	greenObject = greenScoreService.SubOne()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(greenObject)
}
