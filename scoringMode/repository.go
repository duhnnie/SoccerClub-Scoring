package scoringMode

import (
	"encoding/json"
	"fmt"
)

type Repository struct {
	items map[string]*ScoringMode
}

func NewRepo(items map[string]*ScoringMode) *Repository {
	return &Repository{items: items}
}

// TODO: remove this or move it to a repository/factory
func NewRepoFromData(data []byte) (*Repository, error) {
	var scoringModesSlice []*ScoringMode
	scoringModeItems := map[string]*ScoringMode{}

	if err := json.Unmarshal(data, &scoringModesSlice); err != nil {
		return nil, err
	}

	for _, scoringModeItem := range scoringModesSlice {
		scoringModeItems[scoringModeItem.ID] = scoringModeItem
	}

	return NewRepo(scoringModeItems), nil
}

func (r *Repository) Get(scoringModeID string) (*ScoringMode, error) {
	if s, exists := r.items[scoringModeID]; exists {
		return s, nil
	}

	return nil, fmt.Errorf("no \"%s\" scoring mode found", scoringModeID)
}

// func (r *Repository) Resolve(scoringModeID string, vars types.VariableContainer, criteria types.ScoringCriteria) ([]*types.PredictionHit, error) {
// 	scoringMode, exists := r.items[scoringModeID]
// 	if !exists {
// 		return nil, ErrorUknownScoringMode(scoringModeID)
// 	}

// 	return scoringMode.Resolve(vars, criteria, r.scoringItemsRepo)
// }
