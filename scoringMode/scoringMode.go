package scoringMode

import (
	"slices"

	"github.com/duhnnie/godash"
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type ScoringMode struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Strategy    []*ScoringStep `json:"strategy"`
}

func (s *ScoringMode) sumAll(scoringItems []string, vars types.VariableContainer, criteria types.ScoringCriteria, scoringItemsRepo *scoring.Repository) ([]*types.PredictionHit, error) {
	// NOTE: panic recovering will be covered as long this method is called through ScoringMode.Resolve,
	// right now there's no problem since this method is private and is only called by ScoringMode.Resolve.
	var scores []*types.PredictionHit

	return godash.Reduce(scoringItems, func(scores []*types.PredictionHit, scoringItem string, _ int, _ []string) []*types.PredictionHit {
		if res, err := scoringItemsRepo.ExecuteItem(scoringItem, vars); err != nil {
			panic(err)
		} else if res {
			points, exists := criteria[scoringItem]

			if !exists {
				panic(ErrorNoPointsForCriteria(scoringItem))
			}

			scores = append(scores, &types.PredictionHit{
				ScoringItem: scoringItem,
				Points:      points,
			})
		}

		return scores
	}, scores), nil
}

func (s *ScoringMode) sumFirstHit(scoringItems []string, vars types.VariableContainer, criteria types.ScoringCriteria, scoringItemsRepo *scoring.Repository) ([]*types.PredictionHit, error) {
	// NOTE: panic recovering will be covered as long this method is called through ScoringMode.Resolve,
	// right now there's no problem since this method is private and is only called by ScoringMode.Resolve.
	hitIndex := slices.IndexFunc(scoringItems, func(scoringItem string) bool {
		res, err := scoringItemsRepo.ExecuteItem(scoringItem, vars)

		if err != nil {
			panic(err)
		}

		return res
	})

	var scores []*types.PredictionHit

	if hitIndex == -1 {
		return scores, nil
	}

	scoringItem := scoringItems[hitIndex]
	points, exists := criteria[scoringItem]

	if !exists {
		panic(ErrorNoPointsForCriteria(scoringItem))
	}

	score := &types.PredictionHit{
		ScoringItem: scoringItem,
		Points:      points,
	}

	scores = append(scores, score)

	return scores, nil
}

func (s *ScoringMode) Resolve(vars types.VariableContainer, criteria types.ScoringCriteria, scoringItemsRepo *scoring.Repository) (res []*types.PredictionHit, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()

	var scores []*types.PredictionHit

	return godash.Reduce(s.Strategy, func(acc []*types.PredictionHit, item *ScoringStep, _ int, _ []*ScoringStep) []*types.PredictionHit {
		if item.SkipIfScore && len(acc) > 0 {
			return acc
		}

		var scores []*types.PredictionHit
		var err error = nil

		switch item.Type {
		case ScoringStepTypeSumAll:
			scores, err = s.sumAll(item.Items, vars, criteria, scoringItemsRepo)
		case ScoringStepTypeSumFirstHit:
			scores, err = s.sumFirstHit(item.Items, vars, criteria, scoringItemsRepo)
		default:
			panic(ErrorUknownScoringStepMode(item.Type))
		}

		if err != nil {
			panic(err)
		}

		return append(acc, scores...)
	}, scores), nil
}
