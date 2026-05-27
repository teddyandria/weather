// Définit le modèle interne unifié, indépendant de tout format source.
package weather

import "time"

type Station struct {
	ID           string
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
	Manufacturer  string
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
