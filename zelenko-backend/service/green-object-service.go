package service

import (
	"fmt"
	"zelenko-backend/dto"
	"zelenko-backend/model"
	"zelenko-backend/repository"

	"github.com/google/uuid"
)

type GreenObjectService interface {
	AddObject(dto.IGreenObject) model.GreenObject
	FindAll() []model.GreenObject
}

type greenObjectService struct{}

var (
	greenObjectRepository repository.GreenObjectRepository
)

func NewGreenObjectService(gor repository.GreenObjectRepository) GreenObjectService {

	greenObjectRepository = gor
	return &greenObjectService{}
}

func (s *greenObjectService) AddObject(request dto.IGreenObject) model.GreenObject {

	loc := model.Location{
		ID:        uuid.New(),
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Street:    request.Street,
		City:      request.City,
		Country:   request.Country,
	}

	gs := model.GreenScore{
		ID:           uuid.New(),
		Verification: 0,
		Report:       0,
		TrashRank:    model.New,
	}

	gObj := model.GreenObject{
		ID:           uuid.New(),
		LocationName: request.LocationName,
		Location:     loc,
		Shape:        request.Shape,
		TrashType:    request.TrashType,
		GreenScore:   gs,
		Disabled:     false,
	}

	newObject := greenObjectRepository.Save(gObj)

	return newObject

}

func (s *greenObjectService) FindAll() []model.GreenObject {
	fmt.Println("Srv")
	var list []model.GreenObject
	var err error
	list, err = greenObjectRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	return list

}
