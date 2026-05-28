// Lecture et conversion des données météo depuis un fichier JSON vers le modèle interne.
package main

import (
	"encoding/json"
	"os"
	"time"
)

type jsonStation struct {
	ID           string             `json:"id"`
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
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	//dateInstalled est en string car dans le json, c'est sous format "2020-01-02" et pas un timestamp.
	// encoding/json ne sait pas parser ça en time.Time directement.
	DateInstalled string `json:"installed_on"`
}

type jsonObservations struct {
	//Celsius
	Temperature float64   `json:"temperature_celsius"`
	Sky         string    `json:"conditions"`
	Wind        jsonWind  `json:"wind"`
	Notes       *string   `json:"notes"`
	Time        time.Time `json:"timestamp"`
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
	st.ID = s.ID

	//convertir les coordonnées pour les stocker dans le model interne Station
	st.Coordinates = Coordinates{
		Latitude:  s.Location.Latitude,
		Longitude: s.Location.Longitude,
		Altitude:  s.Altitude,
	}

	//pays
	st.Country = countryISO[s.Country]

	//device
	//parsing de la date d'installation du device, qui est au format "2020-01-02" dans le JSON, en time.Time
	installedOn, err := time.Parse("2006-01-02", s.Device.DateInstalled)
	if err != nil {
		installedOn = time.Time{} // valeur par défaut si le parsing échoue
	}

	st.Device = Device{
		Manufacturer:  s.Device.Manufacturer,
		Type:          s.Device.Type,
		DateInstalled: installedOn,
	}

	//observations
	for _, o := range s.Observations {
		st.Observations = append(st.Observations, Observations{
			Temperature: o.Temperature,
			Sky:         o.Sky,
			Wind: Wind{
				Speed:     o.Wind.Speed,
				Direction: o.Wind.Direction,
			},
			Notes: o.Notes,
			Time:  o.Time,
		})
	}

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
