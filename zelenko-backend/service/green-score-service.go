package service

import (
	"zelenko-backend/crdt"
	"zelenko-backend/dummy"
	"zelenko-backend/model"
	"zelenko-backend/repository"
)

type GreenScoreService interface {
	AddOne() model.GreenObject
	SubOne() model.GreenObject
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
// TODO: Must implement code that from time to time writes from redis to relational database to keep data
// Test for redis cli: redis-cli HGET dfd16f29-0d34-4e42-8ce4-f9e476d9a277 Verification
func (s *greenScoreService) AddOne() model.GreenObject {

	var g_count int
	g_counter.Increment(dummyObject.ID.String())
	g_count = g_counter.GetValue(dummyObject.ID.String())

	dummyObject.GreenScore.Verification = g_count

	dummyObject = greenScoreRepository.Change(dummyObject)
	//fmt.Println(greenScoreRepository.GetAttributeForObject(dummyObject.ID.String(), "Verification"))
	return dummyObject

}

func (s *greenScoreService) SubOne() model.GreenObject {

	var g_count int
	g_counter.Decrement(dummyObject.ID.String())
	g_count = g_counter.GetValue(dummyObject.ID.String())
	dummyObject.GreenScore.Verification = g_count

	dummyObject = greenScoreRepository.Change(dummyObject)
	//fmt.Println(greenScoreRepository.GetAttributeForObject(dummyObject.ID.String(), "Verification"))
	return dummyObject

}
