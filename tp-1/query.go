// Fonctions de requête sur une liste de stations chargées en mémoire.
package weather

// Renvoie les stations d'un pays donné (code ISO).
func FilterByCountry(stations []Station, iso string) []Station {
	var result []Station
	for _, s := range stations {
		if s.Country == iso {
			result = append(result, s)
		}
	}
	return result
}

// Renvoie la station avec la rafale de vent la plus forte sur l'ensemble du dataset, avec la valeur correspondante.
func MaxWindGust(stations []Station) (Station, float64) {
	var maxStation Station
	var maxSpeed float64
	for _, s := range stations {
		for _, obs := range s.Observations {
			if obs.Wind.Speed > maxSpeed {
				maxSpeed = obs.Wind.Speed
				maxStation = s
			}
		}
	}
	return maxStation, maxSpeed
}

// Renvoie la température moyenne (°C) sur l'ensemble des observations d'une station.
func AvgTemperature(s Station) float64 {
	if len(s.Observations) == 0 {
		return 0
	}
	var total float64
	for _, obs := range s.Observations {
		total += obs.Temperature
	}
	return total / float64(len(s.Observations))
}
