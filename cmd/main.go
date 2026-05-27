package main

import (
	"fmt"

	"github.com/efrei/weather"
)

func main() {
	jsonStations, err := weather.LoadFromJSON("weather_data.json")
	if err != nil {
		fmt.Println("Erreur JSON:", err)
		return
	}

	xmlStations, err := weather.LoadFromXML("weather_data.xml")
	if err != nil {
		fmt.Println("Erreur XML:", err)
		return
	}

	// Comparaison JSON vs XML
	fmt.Println("Stations JSON: %d | Stations XML: %d\n", len(jsonStations), len(xmlStations))
	jsonObs := 0
	for _, s := range jsonStations {
		jsonObs += len(s.Observations)
	}
	xmlObs := 0
	for _, s := range xmlStations {
		xmlObs += len(s.Observations)
	}
	fmt.Println("Observations JSON: %d | Observations XML: %d\n", jsonObs, xmlObs)

	// Station la plus ventée
	station, gust := weather.MaxWindGust(jsonStations)
	fmt.Println("Station la plus ventée: %s (%.1f km/h)\n", station.ID, gust)

	// Température moyenne de Bordeaux
	for _, s := range jsonStations {
		if s.ID == "FR-BOR-001" {
			fmt.Printf("Température moyenne Bordeaux Mérignac: %.2f °C\n", weather.AvgTemperature(s))
		}
	}
}
