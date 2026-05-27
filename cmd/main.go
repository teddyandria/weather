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

	jsonObs := 0
	for _, s := range jsonStations {
		jsonObs += len(s.Observations)
	}
	xmlObs := 0
	for _, s := range xmlStations {
		xmlObs += len(s.Observations)
	}

	fmt.Printf("JSON : %d stations, %d observations\n", len(jsonStations), jsonObs)
	fmt.Printf("XML  : %d stations, %d observations\n", len(xmlStations), xmlObs)

	coherence := "OK"
	if len(jsonStations) != len(xmlStations) || jsonObs != xmlObs {
		coherence = "KO"
	}
	fmt.Printf("Cohérence : %s\n", coherence)

	station, gust := weather.MaxWindGust(jsonStations)
	fmt.Printf("Station la plus ventée : %s (%.1f km/h)\n", station.ID, gust)

	for _, s := range jsonStations {
		if s.ID == "FR-BOR-001" {
			fmt.Printf("Temp. moyenne Bordeaux Mérignac : %.1f °C\n", weather.AvgTemperature(s))
			break
		}
	}
}
