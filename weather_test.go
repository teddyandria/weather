package weather

import (
	"fmt"
	"testing"
)

// test permettant de vérifier que le chargement du JSON fonctionne correctement
func TestLoadJSON(t *testing.T) {
	stations, err := LoadFromJSON("weather_data.json")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Loaded %d stations\n : ", len(stations))
	fmt.Println("Observations : ", len(stations[0].Observations))
}
