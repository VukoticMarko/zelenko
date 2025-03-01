package repository

import (
	"database/sql"
	"fmt"
	"zelenko-backend/model"

	_ "github.com/lib/pq"
)

type GreenObjectRepository interface {
	Save(user model.GreenObject) model.GreenObject
	FindAll() ([]model.GreenObject, error)
	UpdateOne(model.GreenObject) model.GreenObject
	FindOne(objectID string) (model.GreenObject, error)
	DeleteOne(model.GreenObject)
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

	sqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostSQL, portSQL, userSQL, passwordSQL, dbnameSQL)

	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := `
	SELECT
		go."Id", go."LocationName", go."Shape", go."TrashType", go."Disabled",
		l."Id" AS "Location.Id", l."Latitude", l."Longitude", l."Street", l."City", l."Country",
		gs."Id" AS "GreenScore.Id", gs."Verification", gs."Report", gs."TrashRank"
	FROM "GreenObject" go
	LEFT JOIN "Location" l ON go."Location" = l."Id"
	LEFT JOIN "GreenScore" gs ON go."GreenScore" = gs."Id"
`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var greenObjects []model.GreenObject

	for rows.Next() {
		var greenObject model.GreenObject
		var location model.Location
		var greenScore model.GreenScore

		err := rows.Scan(
			&greenObject.ID, &greenObject.LocationName, &greenObject.Shape, &greenObject.TrashType, &greenObject.Disabled,
			&location.ID, &location.Latitude, &location.Longitude, &location.Street, &location.City, &location.Country,
			&greenScore.ID, &greenScore.Verification, &greenScore.Report, &greenScore.TrashRank,
		)

		if err != nil {
			return nil, err
		}

		greenObject.Location = location
		greenObject.GreenScore = greenScore

		greenObjects = append(greenObjects, greenObject)
	}

	return greenObjects, nil
}

func (*goRepository) FindOne(objectID string) (model.GreenObject, error) {

	var result model.GreenObject

	sqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostSQL, portSQL, userSQL, passwordSQL, dbnameSQL)

	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		return result, err
	}

	defer db.Close()
	query := `
	SELECT
		go."Id", go."LocationName", go."Shape", go."TrashType", go."Disabled",
		l."Id" AS "Location.Id", l."Latitude", l."Longitude", l."Street", l."City", l."Country",
		gs."Id" AS "GreenScore.Id", gs."Verification", gs."Report", gs."TrashRank"
	FROM "GreenObject" go
	LEFT JOIN "Location" l ON go."Location" = l."Id"
	LEFT JOIN "GreenScore" gs ON go."GreenScore" = gs."Id"
	WHERE go."Id" = $1
`

	var greenObject model.GreenObject
	var location model.Location
	var greenScore model.GreenScore

	err = db.QueryRow(query, objectID).Scan(
		&greenObject.ID, &greenObject.LocationName, &greenObject.Shape, &greenObject.TrashType, &greenObject.Disabled,
		&location.ID, &location.Latitude, &location.Longitude, &location.Street, &location.City, &location.Country,
		&greenScore.ID, &greenScore.Verification, &greenScore.Report, &greenScore.TrashRank,
	)

	if err != nil {
		return result, err
	}

	greenObject.Location = location
	greenObject.GreenScore = greenScore

	result = greenObject

	return result, nil
}

func (*goRepository) UpdateOne(greenObject model.GreenObject) model.GreenObject {

	sqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostSQL, portSQL, userSQL, passwordSQL, dbnameSQL)

	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	greenObjectQuery := `
    UPDATE "GreenObject"
    SET
        "LocationName" = $1,
        "Shape" = $2,
        "TrashType" = $3,
        "Disabled" = $4
    WHERE "Id" = $5
    `

	locationQuery := `
    UPDATE "Location"
    SET
        "Latitude" = $1,
        "Longitude" = $2,
        "Street" = $3,
        "City" = $4,
        "Country" = $5
    WHERE "Id" = $6
    `

	greenScoreQuery := `
    UPDATE "GreenScore"
    SET
        "Verification" = $1,
        "Report" = $2,
        "TrashRank" = $3
    WHERE "Id" = $4
    `

	_, err = db.Exec(greenObjectQuery,
		greenObject.LocationName,
		greenObject.Shape,
		greenObject.TrashType,
		greenObject.Disabled,
		greenObject.ID,
	)

	_, err = db.Exec(locationQuery,
		greenObject.Location.Latitude,
		greenObject.Location.Longitude,
		greenObject.Location.Street,
		greenObject.Location.City,
		greenObject.Location.Country,
		greenObject.Location.ID,
	)

	_, err = db.Exec(greenScoreQuery,
		greenObject.GreenScore.Verification,
		greenObject.GreenScore.Report,
		greenObject.GreenScore.TrashRank,
		greenObject.GreenScore.ID,
	)

	return greenObject
}

func (*goRepository) DeleteOne(greenObject model.GreenObject) {

	sqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostSQL, portSQL, userSQL, passwordSQL, dbnameSQL)

	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {

	}

	queryDeleteLocation := `DELETE FROM "GreenObject" WHERE "Location" = $1`
	_, err = tx.Exec(queryDeleteLocation, greenObject.ID)
	if err != nil {
		tx.Rollback()

	}

	queryDeleteScore := `DELETE FROM "GreenObject" WHERE "GreenScore" = $1`
	_, err = tx.Exec(queryDeleteScore, greenObject.ID)
	if err != nil {
		tx.Rollback()

	}

	queryDeleteObject := `DELETE FROM "GreenObject" WHERE "Id" = $1`
	_, err = tx.Exec(queryDeleteObject, greenObject.ID)
	if err != nil {
		tx.Rollback()

	}
}
