package main

type Store struct {
	stations map[string]Station
}

func NewStore() *Store {
	return &Store{
		stations: make(map[string]Station),
	}
}

func (s *Store) Put(st Station) {
	s.stations[st.ID] = st
}

func (s *Store) Has(id string) bool {
	_, exists := s.stations[id]
	return exists
}

func (s *Store) Get(id string) (Station, bool) {
	st, exists := s.stations[id]
	return st, exists
}

func (s *Store) Delete(id string) bool {
	if _, exists := s.stations[id]; exists {
		delete(s.stations, id)
		return true
	}
	return false
}

func (s *Store) All() []Station {
	stations := make([]Station, 0, len(s.stations))
	for _, st := range s.stations {
		stations = append(stations, st)
	}
	return stations
}
