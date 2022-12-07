package model

//easyjson:json
type Response []struct {
	ID           string `json:"id"`
	Dislocation  string `json:"dislocation"`
	TrafficSigns string `json:"traffic_signs"`
	Street       string `json:"street,omitempty"`
	Seats        string `json:"seats"`
	Stopping     string `json:"stopping"`
	ParkingBay   string `json:"parking_bay"`
	Avenue       string `json:"avenue,omitempty"`
}
