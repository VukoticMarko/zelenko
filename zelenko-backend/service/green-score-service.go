package service

import (
	"fmt"
	"zelenko-backend/crdt"
	"zelenko-backend/dummy"
	"zelenko-backend/model"
	"zelenko-backend/repository"
)

type GreenScoreService interface {
	AddOne() model.GreenObject
	SubOne()
}

type greenScoreService struct{}

var (
	dmmy                 dummy.Dummy
	dummyObject          model.GreenObject
	g_counter            crdt.GCounter
	greenScoreRepository repository.GreenScoreRepository
)

func NewGreenScoreService(gsr repository.GreenScoreRepository, gc crdt.GCounter) GreenScoreService {

	g_counter = gc
	greenScoreRepository = gsr
	dummyObject = dmmy.GetDummy1()
	return &greenScoreService{}
}

// TODO: Later in app development everything that is dummy object to replace with real dynamic object
func (s *greenScoreService) AddOne() model.GreenObject {

	var g_count int
	g_counter.Increment(dummyObject.ID.String())
	fmt.Println("Init")
	fmt.Println(g_count)
	g_count = g_counter.GetValue(dummyObject.ID.String())

	dummyObject.GreenScore.Verification = g_count

	dummyObject = greenScoreRepository.AddOne(dummyObject)
	fmt.Println("Final")
	fmt.Println(dummyObject.GreenScore.Verification)
	return dummyObject
}

func (s *greenScoreService) SubOne() {}
