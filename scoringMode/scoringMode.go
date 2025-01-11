package scoringMode

import (
	"errors"

	"github.com/duhnnie/godash"
	"github.com/duhnnie/soccerclub-scoring/constraint"
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type ScoringMode struct {
	// TODO: remove json fields
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Strategy    []ScoringStrategy `json:"strategy"`
	// TODO: implement constraints
	Constraints []constraint.Constraint
}

func (s *ScoringMode) AreConstraintsMeet(context *types.ScoringCriteria) bool {
	return godash.Every(s.Constraints, func(c constraint.Constraint, _ int, _ []constraint.Constraint) bool {
		return c.IsMet(context)
	})
}

func (s *ScoringMode) Resolve(context types.PredictionContext, predictions []*types.Prediction, criteria *types.ScoringCriteria, scoringItems types.Store[scoring.Item]) (res types.PredictionScores, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()

	if !s.AreConstraintsMeet(criteria) {
		return nil, errors.New("constraints are not met")
	}

	return godash.Reduce(predictions, func(acc types.PredictionScores, p *types.Prediction, _ int, _ []*types.Prediction) types.PredictionScores {
		context.SetPrediction(p)

		hits := godash.Reduce(s.Strategy, func(acc []*types.PredictionHit, item ScoringStrategy, _ int, _ []ScoringStrategy) []*types.PredictionHit {
			if item.SkipIfHit() && len(acc) > 0 {
				return acc
			}

			scores, err := item.Execute(scoringItems, context, criteria)

			if err != nil {
				panic(err)
			}

			return append(acc, scores...)
		}, []*types.PredictionHit{})

		acc[p.GetID()] = hits

		return acc
	}, types.PredictionScores{}), nil
}
