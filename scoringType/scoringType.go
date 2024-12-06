package scoringType

import (
	"slices"

	"github.com/duhnnie/soccerclub-scoring/arrayHelpers"
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/variable"
)

type ScoringCriteria = map[string]float64

type ScoringType struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Strategy    []*ScoringStep `json:"strategy"`
}

func (s *ScoringType) sumAll(scoringItems []string, vars *variable.Repository, criteria ScoringCriteria, scoringItemsRepo *scoring.Repository) ([]*scoring.Score, error) {
	var scores []*scoring.Score

	return arrayHelpers.Reduce(scoringItems, func(scores []*scoring.Score, scoringItem string, _ int) []*scoring.Score {
		// TODO: recover from panic
		if res, err := scoringItemsRepo.ExecuteItem(scoringItem, vars); err != nil {
			panic(err)
		} else if res {
			points, exists := criteria[scoringItem]

			if !exists {
				panic(ErrorNoPointsForSCriteria(scoringItem))
			}

			scores = append(scores, &scoring.Score{
				ScoringItem: scoringItem,
				Points:      int(points),
			})
		}

		return scores
	}, scores), nil
}

func (s *ScoringType) sumFirstHit(scoringItems []string, vars *variable.Repository, criteria ScoringCriteria, scoringItemsRepo *scoring.Repository) ([]*scoring.Score, error) {
	// TODO: recover from panic
	hitIndex := slices.IndexFunc(scoringItems, func(scoringItem string) bool {
		res, err := scoringItemsRepo.ExecuteItem(scoringItem, vars)

		if err != nil {
			panic(err)
		}

		return res
	})

	var scores []*scoring.Score

	if hitIndex == -1 {
		return scores, nil
	}

	scoringItem := scoringItems[hitIndex]
	points, exists := criteria[scoringItem]

	if !exists {
		panic(ErrorNoPointsForSCriteria(scoringItem))
	}

	score := &scoring.Score{
		ScoringItem: scoringItem,
		Points:      int(points),
	}

	scores = append(scores, score)

	return scores, nil
}

func (s *ScoringType) Resolve(vars *variable.Repository, criteria ScoringCriteria, scoringItemsRepo *scoring.Repository) ([]*scoring.Score, error) {
	var scores []*scoring.Score

	// TODO: recover from panic
	return arrayHelpers.Reduce(s.Strategy, func(acc []*scoring.Score, item *ScoringStep, _ int) []*scoring.Score {
		if item.SkipIfScore && len(acc) > 0 {
			return acc
		}

		var scores []*scoring.Score
		var err error = nil

		switch item.Type {
		case ScoringStepTypeSumAll:
			scores, err = s.sumAll(item.Items, vars, criteria, scoringItemsRepo)
		case ScoringStepTypeSumFirstHit:
			scores, err = s.sumFirstHit(item.Items, vars, criteria, scoringItemsRepo)
		default:
			panic(ErrorUknownScoringStepType(item.Type))
		}

		if err != nil {
			panic(err)
		}

		return append(acc, scores...)
	}, scores), nil
}
