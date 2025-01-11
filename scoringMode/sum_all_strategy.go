package scoringMode

import (
	"github.com/duhnnie/godash"
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type SumAllStrategy struct {
	skipIfHit    bool
	scoringItems []string
}

func NewSumAllStrategy(skipIfHit bool, scoringItems []string) *SumAllStrategy {
	return &SumAllStrategy{skipIfHit, scoringItems}
}

func (s *SumAllStrategy) SkipIfHit() bool {
	return s.skipIfHit
}

func (s *SumAllStrategy) Execute(scoringItemsStore types.Store[scoring.Item], context types.VariableContainer, criteria *types.ScoringCriteria) (res []*types.PredictionHit, err error) {
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

	return godash.Reduce(s.scoringItems, func(hits []*types.PredictionHit, scoringItemId string, _ int, _ []string) []*types.PredictionHit {
		scoringItem, found := scoringItemsStore.Get(scoringItemId)

		if !found {
			panic(scoring.ScoringItemNotFoundError(scoringItemId))
		}

		if res, err := scoringItem.Resolve(context); err != nil {
			panic(err)
		} else if !res {
			return hits
		} else if score, err := criteria.GetScore(scoringItemId); err != nil {
			panic(err)
		} else {
			hit := &types.PredictionHit{ScoringItem: scoringItemId, Points: score}
			hits = append(hits, hit)
		}

		return hits
	}, []*types.PredictionHit{}), nil
}
