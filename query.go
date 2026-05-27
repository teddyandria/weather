package weather

func FilterByCountry(stations []Station, iso string) []Station {
	var result []Station
	for _, s := range stations {
		if s.Country == iso {
			result = append(result, s)
		}
	}
	return result
}

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
