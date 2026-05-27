// Lecture et conversion des données météo depuis un fichier XML vers le modèle interne.
package weather

import (
	"encoding/xml"
	"os"
	"strconv"
	"time"
)

type xmlDataset struct {
	Stations []xmlStation `xml:"station"`
}

type xmlStation struct {
	Country      string           `xml:"country,attr"`
	Coordinates  xmlCoordinates   `xml:"coordinates"`
	Device       xmlDevice        `xml:"device"`
	Observations []xmlObservation `xml:"observations>observation"`
}

type xmlCoordinates struct {
	Lat      float64 `xml:"lat,attr"`
	Lon      float64 `xml:"lon,attr"`
	Altitude int     `xml:"altitude,attr"`
}

type xmlDevice struct {
	Vendor string `xml:"vendor,attr"`
	Model  string `xml:"model,attr"`
	Since  string `xml:"since,attr"`
}

type xmlObservation struct {
	At       string       `xml:"at,attr"`
	Sky      string       `xml:"sky,attr"`
	Measures []xmlMeasure `xml:"measure"`
	Wind     xmlWind      `xml:"wind"`
	Note     *string      `xml:"note"`
}

type xmlMeasure struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type xmlWind struct {
	Speed     string `xml:"speed,attr"`
	Direction string `xml:"direction,attr"`
}

func (s xmlStation) convert() (st Station) {
	installedOn, err := time.Parse("2006-01-02", s.Device.Since)
	if err != nil {
		installedOn = time.Time{}
	}

	st = Station{
		Country: s.Country,
		Coordinates: Coordinates{
			Latitude:  s.Coordinates.Lat,
			Longitude: s.Coordinates.Lon,
			Altitude:  s.Coordinates.Altitude,
		},
		Device: Device{
			Manufacturer:  s.Device.Model,
			Type:          s.Device.Vendor,
			DateInstalled: installedOn,
		},
	}

	for _, obs := range s.Observations {
		t, _ := time.Parse(time.RFC3339, obs.At)

		var temp float64
		for _, m := range obs.Measures {
			if m.Type == "temperature" {
				temp, _ = strconv.ParseFloat(m.Value, 64)
				break
			}
		}

		speed, _ := strconv.ParseFloat(obs.Wind.Speed, 64)
		//même pruncipe que strconv.ParseFloat, mais pour les int. atoi = ascii to int.
		direction, _ := strconv.Atoi(obs.Wind.Direction)

		st.Observations = append(st.Observations, Observations{
			Temperature: temp,
			Sky:         obs.Sky,
			Wind:        Wind{Speed: speed, Direction: direction},
			Notes:       obs.Note,
			Time:        t,
		})
	}

	return st
}

func LoadFromXML(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dataset xmlDataset
	err = xml.Unmarshal(data, &dataset)
	if err != nil {
		return nil, err
	}

	var result []Station
	for _, s := range dataset.Stations {
		result = append(result, s.convert())
	}
	return result, nil
}
