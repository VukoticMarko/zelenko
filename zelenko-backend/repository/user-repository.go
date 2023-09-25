package repository

import (
	"database/sql"
	"fmt"
	"zelenko-backend/model"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	Save(user model.User) model.User
	FindUsers() ([]model.User, error)
	FindOne(id uuid.UUID) model.User
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (*userRepository) Save(user model.User) model.User {

	// Connection string for db
	sqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostSQL, portSQL, userSQL, passwordSQL, dbnameSQL)

	// Open db
	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	// Insert into db
	var id = user.ID.String()
	var urID = user.UserRank.ID.String()
	// Without picture and bday

	insert1 := `insert into "UserRank"("Id", "UserPoints", "UserRank") values ($1, $2, $3)`
	_, err = db.Exec(insert1, urID, user.UserRank.UserPoints, user.UserRank.UserRank)
	if err != nil {
		panic(err)
	}

	insert2 := `insert into "User"("Id", "Username", "Mail", "Password", "Name", "Surname", "City",
                      "Country", "Sex", "Disabled", "UserRank", "Role") values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err = db.Exec(insert2, id, user.Username, user.Mail, user.Password, user.Name, user.Surname, user.City,
		user.Country, user.Sex, user.Disabled, urID, user.Role)
	if err != nil {
		panic(err)
	}

	return user
}

func (*userRepository) FindUsers() ([]model.User, error) {
	panic("not implemented")
}

func (*userRepository) FindOne(id uuid.UUID) model.User {
	panic("not implemented")
}
