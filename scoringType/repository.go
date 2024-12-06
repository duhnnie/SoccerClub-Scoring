package scoringType

import (
	"encoding/json"

	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/variable"
)

type Repository struct {
	items            map[string]*ScoringType
	scoringItemsRepo *scoring.Repository
}

func NewRepo(items map[string]*ScoringType, scoringItemsRepo *scoring.Repository) *Repository {
	return &Repository{
		items:            items,
		scoringItemsRepo: scoringItemsRepo,
	}
}

func NewRepoFromData(data []byte, scoringItemsRepo *scoring.Repository) (*Repository, error) {
	var scoringTypesSlice []*ScoringType
	scoringTypeItems := map[string]*ScoringType{}

	if err := json.Unmarshal(data, &scoringTypesSlice); err != nil {
		return nil, err
	}

	for _, scoringTypeItem := range scoringTypesSlice {
		scoringTypeItems[scoringTypeItem.ID] = scoringTypeItem
	}

	return NewRepo(scoringTypeItems, scoringItemsRepo), nil
}

func (r *Repository) Resolve(scoringTypeID string, vars *variable.Repository, criteria ScoringCriteria) ([]*scoring.Score, error) {
	scoringType, exists := r.items[scoringTypeID]
	if !exists {
		return nil, ErrorUknownScoringType(scoringTypeID)
	}

	return scoringType.Resolve(vars, criteria, r.scoringItemsRepo)
}
