package scoringMode

import (
	"encoding/json"

	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type Repository struct {
	items            map[string]*ScoringMode
	scoringItemsRepo *scoring.Repository
}

func NewRepo(items map[string]*ScoringMode, scoringItemsRepo *scoring.Repository) *Repository {
	return &Repository{
		items:            items,
		scoringItemsRepo: scoringItemsRepo,
	}
}

func NewRepoFromData(data []byte, scoringItemsRepo *scoring.Repository) (*Repository, error) {
	var scoringModesSlice []*ScoringMode
	scoringModeItems := map[string]*ScoringMode{}

	if err := json.Unmarshal(data, &scoringModesSlice); err != nil {
		return nil, err
	}

	for _, scoringModeItem := range scoringModesSlice {
		scoringModeItems[scoringModeItem.ID] = scoringModeItem
	}

	return NewRepo(scoringModeItems, scoringItemsRepo), nil
}

func (r *Repository) Resolve(scoringModeID string, vars types.VariableContainer, criteria types.ScoringCriteria) ([]*types.PredictionHit, error) {
	scoringMode, exists := r.items[scoringModeID]
	if !exists {
		return nil, ErrorUknownScoringMode(scoringModeID)
	}

	return scoringMode.Resolve(vars, criteria, r.scoringItemsRepo)
}
