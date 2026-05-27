package weather

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonStation struct {
	Country      string             `json:"country"`
	Altitude     int                `json:"altitude_m"`
	Location     jsonCoordinates    `json:"location"`
	Observations []jsonObservations `json:"observations"`
	Device       jsonDevice         `json:"device"`
}

type jsonCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type jsonWind struct {
	Speed     float64 `json:"speed_kmh"`
	Direction int     `json:"direction_deg"`
}

type jsonDevice struct {
	Manufacturer  string `json:"manufacturer"`
	Type          string `json:"type"`
	DateInstalled string `json:"installed_on"`
}

type jsonObservations struct {
	//Celsius
	Temperature float64  `json:"temperature_celsius"`
	Sky         string   `json:"conditions"`
	Wind        jsonWind `json:"wind"`
	Notes       *string  `json:"notes"`
	Time        string   `json:"timestamp"`
}

// func de conversion valeur country du JSON "France" en "FR" -> iso 2 lettres
func (s jsonStation) convert() (st Station) {
	fmt.Println()
	return
}
func LoadFromJSON(path string) ([]Station, error) {
	//lire un fichier avec os.ReadFile
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}
	//décoder le json dans []jsonStation avec json.Unmarshal
	var stations []jsonStation
	err = json.Unmarshal(data, &stations)

	if err != nil {
		return nil, err
	}

	//convertir []jsonStation en []Station
	var result []Station
	for i, s := range stations {
		// Conversion logique ici si nécessaire
		result[i] = s.convert()
	}
	return result, nil
}
