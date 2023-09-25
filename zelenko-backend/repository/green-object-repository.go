package repository

import (
	"database/sql"
	"fmt"
	"zelenko-backend/model"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type GreenObjectRepository interface {
	Save(user model.GreenObject) model.GreenObject
	FindAll() ([]model.GreenObject, error)
	FindOne(id uuid.UUID) model.GreenObject
}

type goRepository struct{}

func NewGreenObjectRepository() GreenObjectRepository {
	return &goRepository{}
}

func (*goRepository) Save(greenObject model.GreenObject) model.GreenObject {

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
	var id = greenObject.ID.String()
	var gsID = greenObject.GreenScore.ID.String()
	var lID = greenObject.Location.ID.String()

	insert1 := `insert into "Location"("Id", "Latitude", "Longitude", "Street", "City", "Country") values ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(insert1, lID, greenObject.Location.Latitude, greenObject.Location.Longitude, greenObject.Location.Street,
		greenObject.Location.City, greenObject.Location.Country)
	if err != nil {
		panic(err)
	}

	insert2 := `insert into "GreenScore"("Id", "Verification", "Report", "TrashRank") values ($1, $2, $3, $4)`
	_, err = db.Exec(insert2, gsID, greenObject.GreenScore.Verification, greenObject.GreenScore.Report, greenObject.GreenScore.TrashRank)
	if err != nil {
		panic(err)
	}

	insert3 := `insert into "GreenObject"("Id", "LocationName", "Location", "Shape", "TrashType", "GreenScore", "Disabled") values ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(insert3, id, greenObject.LocationName, lID, greenObject.Shape, greenObject.TrashType,
		gsID, greenObject.Disabled)
	if err != nil {
		panic(err)
	}

	return greenObject
}

func (*goRepository) FindAll() ([]model.GreenObject, error) {
	panic("not implemented")
}

func (*goRepository) FindOne(id uuid.UUID) model.GreenObject {
	panic("not implemented")
}
