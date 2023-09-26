package service

import (
	"zelenko-backend/crdt"
	"zelenko-backend/model"
	"zelenko-backend/repository"
)

type GreenScoreService interface {
	AddOne(model.GreenObject) model.GreenObject
	SubOne(model.GreenObject) model.GreenObject
}

type greenScoreService struct{}

var (
	g_counter            crdt.GCounter
	greenScoreRepository repository.GreenScoreRepository
)

func NewGreenScoreService(gsr repository.GreenScoreRepository, gc crdt.GCounter) GreenScoreService {

	g_counter = gc
	greenScoreRepository = gsr
	return &greenScoreService{}
}

// TODO: Must implement code that from time to time writes from redis to relational database to keep data
// Test for redis cli: redis-cli HGET cde5ac3c-5c0f-489e-9342-ba2215037fa5 Verification
func (s *greenScoreService) AddOne(request model.GreenObject) model.GreenObject {

	var g_count int
	g_counter.Increment(request.ID.String())
	g_count = g_counter.GetValue(request.ID.String())

	request.GreenScore.Verification = g_count

	request = greenScoreRepository.Change(request)

	return request
}

func (s *greenScoreService) SubOne(request model.GreenObject) model.GreenObject {

	var g_count int
	g_counter.Decrement(request.ID.String())
	g_count = g_counter.GetValue(request.ID.String())

	request.GreenScore.Verification = g_count

	request = greenScoreRepository.Change(request)

	return request
}
