package dummy

import (
	"zelenko-backend/model"

	"github.com/google/uuid"
)

type Dummy struct {
}

func (Dummy) GetDummy1() model.GreenObject {

	UUID := "dfd16f29-0d34-4e42-8ce4-f9e476d9a277"
	parsedUUID, _ := uuid.Parse(UUID)
	//newUUID := uuid.New()

	gs := model.GreenScore{
		//Verification: 0,
		Report:    5,
		TrashRank: "new",
	}

	dummy := model.GreenObject{
		ID:           parsedUUID,
		LocationName: "Trg Studenata, Novi Sad",
		Shape:        model.TrashCan,
		TrashType:    model.All,
		GreenScore:   gs,
		Disabled:     false,
	}

	return dummy
}
