package store

import (
	"github.com/duhnnie/soccerclub-scoring/repo"
	"github.com/duhnnie/soccerclub-scoring/scoringMode"
)

type ScoringModeStore struct {
	itemsMap map[string]*scoringMode.ScoringMode
}

func NewScoringModeStore() (*ScoringModeStore, error) {
	items, err := repo.GetAll()

	if err != nil {
		return nil, err
	}

	itemsMap := map[string]*scoringMode.ScoringMode{}

	for _, i := range items {
		itemsMap[i.ID] = i
	}

	return &ScoringModeStore{itemsMap: itemsMap}, nil
}

func (s *ScoringModeStore) Get(id string) (*scoringMode.ScoringMode, bool) {
	if v, exists := s.itemsMap[id]; !exists {
		return v, false
	} else {
		return v, true
	}
}
