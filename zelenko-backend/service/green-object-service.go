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
	gsRepository          repository.GreenScoreRepository
	counter               int
)

func NewGreenObjectService(gor repository.GreenObjectRepository, gsr repository.GreenScoreRepository) GreenObjectService {
	greenScoreRepository = gsr
	greenObjectRepository = gor
	counter = 0
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

	var list []model.GreenObject
	var err error
	list, err = greenObjectRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	UpdateScores(list, 1)

	counter++

	if counter >= 10 {
		UpdateScores(list, 2)
		counter = 0
	}

	return list

}

func UpdateScores(list []model.GreenObject, flag int) []model.GreenObject {
	if flag == 1 {
		for _, element := range list {
			score := gsRepository.GetAttributeForObject(element.ID.String(), "Verification")
			element.GreenScore.Verification = int(score)
		}
		return list
	}
	if flag == 2 {
		for _, element := range list {
			score := gsRepository.GetAttributeForObject(element.ID.String(), "Verification")
			element.GreenScore.Verification = int(score)
			element = greenObjectRepository.UpdateOne(element)
		}
		return list
	}
	return list
}
