package scoringMode

import (
	"slices"

	"github.com/duhnnie/soccerclub-scoring/arrayHelpers"
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/variable"
)

type ScoringCriteria = map[string]float64

type ScoringMode struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Strategy    []*ScoringStep `json:"strategy"`
}

func (s *ScoringMode) sumAll(scoringItems []string, vars *variable.Repository, criteria ScoringCriteria, scoringItemsRepo *scoring.Repository) ([]*scoring.PredictionHit, error) {
	// NOTE: panic recovering will be covered as long this method is called through ScoringMode.Resolve,
	// right now there's no problem since this method is private and is only called by ScoringMode.Resolve.
	var scores []*scoring.PredictionHit

	return arrayHelpers.Reduce(scoringItems, func(scores []*scoring.PredictionHit, scoringItem string, _ int) []*scoring.PredictionHit {
		if res, err := scoringItemsRepo.ExecuteItem(scoringItem, vars); err != nil {
			panic(err)
		} else if res {
			points, exists := criteria[scoringItem]

			if !exists {
				panic(ErrorNoPointsForCriteria(scoringItem))
			}

			scores = append(scores, &scoring.PredictionHit{
				ScoringItem: scoringItem,
				Points:      int(points),
			})
		}

		return scores
	}, scores), nil
}

func (s *ScoringMode) sumFirstHit(scoringItems []string, vars *variable.Repository, criteria ScoringCriteria, scoringItemsRepo *scoring.Repository) ([]*scoring.PredictionHit, error) {
	// NOTE: panic recovering will be covered as long this method is called through ScoringMode.Resolve,
	// right now there's no problem since this method is private and is only called by ScoringMode.Resolve.
	hitIndex := slices.IndexFunc(scoringItems, func(scoringItem string) bool {
		res, err := scoringItemsRepo.ExecuteItem(scoringItem, vars)

		if err != nil {
			panic(err)
		}

		return res
	})

	var scores []*scoring.PredictionHit

	if hitIndex == -1 {
		return scores, nil
	}

	scoringItem := scoringItems[hitIndex]
	points, exists := criteria[scoringItem]

	if !exists {
		panic(ErrorNoPointsForCriteria(scoringItem))
	}

	score := &scoring.PredictionHit{
		ScoringItem: scoringItem,
		Points:      int(points),
	}

	scores = append(scores, score)

	return scores, nil
}

func (s *ScoringMode) Resolve(vars *variable.Repository, criteria ScoringCriteria, scoringItemsRepo *scoring.Repository) (res []*scoring.PredictionHit, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()

	var scores []*scoring.PredictionHit

	return arrayHelpers.Reduce(s.Strategy, func(acc []*scoring.PredictionHit, item *ScoringStep, _ int) []*scoring.PredictionHit {
		if item.SkipIfScore && len(acc) > 0 {
			return acc
		}

		var scores []*scoring.PredictionHit
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
