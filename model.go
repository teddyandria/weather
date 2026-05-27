package weather

import "time"

type Station struct {
	Country      string
	Coordinates  Coordinates
	Observations []Observations
	Device       Device
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
	Altitude  int
}

type Wind struct {
	Speed     float64 //km/h
	Direction int
}

type Device struct {
	Model         string
	Type          string
	DateInstalled time.Time
}

type Observations struct {
	Temperature float64 //Celsius
	Sky         string
	Wind        Wind
	Notes       *string
	Time        time.Time
}
