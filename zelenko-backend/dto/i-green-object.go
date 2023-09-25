package dto

type IGreenObject struct {
	LocationName string
	Shape        string
	TrashType    string

	Latitude  float32
	Longitude float32
	Street    string
	City      string
	Country   string
}
