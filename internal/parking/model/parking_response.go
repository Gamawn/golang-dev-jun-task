package model

import (
	"github.com/mailru/easyjson"
)

//easyjson:json
type Response []Parking

//easyjson:json
type Parking struct {
	ID           string `json:"id"`
	Dislocation  string `json:"dislocation"`
	TrafficSigns string `json:"traffic_signs"`
	Street       string `json:"street,omitempty"`
	Seats        string `json:"seats"`
	Stopping     string `json:"stopping"`
	ParkingBay   string `json:"parking_bay"`
	Avenue       string `json:"avenue,omitempty"`
}

func (r Parking) MarshalBinary() (data []byte, err error) {
	bytes, err := easyjson.Marshal(r)
	return bytes, err
}
