package weather

import (
	"encoding/json"
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

type jsonRoot struct {
	Stations []jsonStation `json:"stations"`
}

var countryISO = map[string]string{
	"France":    "FR",
	"Allemagne": "DE",
	"Espagne":   "ES",
	"Italie":    "IT",
	"Portugal":  "PT",
	"Belgique":  "BE",
	"Suisse":    "CH",
	"Autriche":  "AT",
	"Pays-Bas":  "NL",
	"Norvège":   "NO",
	"Suède":     "SE",
	"Danemark":  "DK",
	"Pologne":   "PL",
	"Tchèque":   "CZ",
}

// func de conversion valeur country du JSON "France" en "FR" -> iso 2 lettres
func (s jsonStation) convert() (st Station) {
	//convertir les coordonnées pour les stocker dans le model interne Station
	st.Coordinates = Coordinates{
		Latitude:  s.Location.Latitude,
		Longitude: s.Location.Longitude,
		Altitude:  s.Altitude,
	}

	//pays
	st.Country = countryISO[s.Country]
	return
}

func LoadFromJSON(path string) ([]Station, error) {
	//lire un fichier avec os.ReadFile
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}
	//décoder le json en utilisant json.Unmarshal via une struct intermédiaire jsonRoot
	var root jsonRoot
	err = json.Unmarshal(data, &root)

	if err != nil {
		return nil, err
	}

	//convertir les jsonStation en Station
	var result []Station
	for _, s := range root.Stations {

		result = append(result, s.convert())
	}
	return result, nil
}
