package service

import (
	"zelenko-backend/model"
	"zelenko-backend/repository"

	"github.com/google/uuid"
)

type UserService interface {
	AddUser(model.User) (model.User, error)
	GetUser(id uuid.UUID) model.User
}

type userService struct{}

var (
	userRepository repository.UserRepository
)

func NewUserService(ur repository.UserRepository) UserService {

	userRepository = ur
	return &userService{}
}

// TODO: Later in app development everything that is dummy object to replace with real dynamic object
// TODO: Must implement code that from time to time writes from redis to relational database to keep data
// Test for redis cli: redis-cli HGET dfd16f29-0d34-4e42-8ce4-f9e476d9a277 Verification
func (s *userService) AddUser(user model.User) (model.User, error) {

	user.ID, _ = uuid.NewUUID()
	newRank := model.UserRank{
		ID:         uuid.New(),
		UserPoints: 0,
		UserRank:   model.Baby,
	}
	user.UserRank = newRank
	user = userRepository.Save(user)
	return user, nil

}

func (s *userService) GetUser(id uuid.UUID) model.User {

	var user model.User

	user = userRepository.FindOne(id)
	return user

}
