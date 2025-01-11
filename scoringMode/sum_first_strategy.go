package scoringMode

import (
	"slices"

	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type SumFirstStrategy struct {
	skipIfHit    bool
	scoringItems []string
}

func NewSumFirstStrategy(skipIfHit bool, scoringItems []string) *SumFirstStrategy {
	return &SumFirstStrategy{skipIfHit, scoringItems}
}

func (s *SumFirstStrategy) SkipIfHit() bool {
	return s.skipIfHit
}

func (s *SumFirstStrategy) Execute(scoringItemsRepo *scoring.Repository, context types.VariableContainer, criteria types.ScoringCriteria) (res []*types.PredictionHit, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				res = nil
				err = e
			} else {
				panic(r)
			}
		}
	}()

	hitIndex := slices.IndexFunc(s.scoringItems, func(scoringItemId string) bool {
		scoringItem := scoringItemsRepo.Get(scoringItemId)

		if scoringItem == nil {
			panic(scoring.ScoringItemNotFoundError(scoringItemId))
		}

		res, err := scoringItem.Resolve(context)

		if err != nil {
			panic(err)
		}

		return res
	})

	var hits []*types.PredictionHit

	if hitIndex == -1 {
		return hits, nil
	}

	scoringItem := s.scoringItems[hitIndex]
	score, err := criteria.GetScore(scoringItem)

	if err != nil {
		return hits, err
	}

	hit := &types.PredictionHit{ScoringItem: scoringItem, Points: score}
	hits = append(hits, hit)

	return hits, nil
}
