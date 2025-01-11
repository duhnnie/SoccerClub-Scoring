package resolver

import (
	"errors"

	bettinggroup "github.com/duhnnie/soccerclub-scoring/bettingGroup"
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/scoringMode"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type PredictionScoreResolver struct {
	scoringItemsRepo  *scoring.Repository
	scoringModesStore types.Store[*scoringMode.ScoringMode]
}

func New(ir *scoring.Repository, mr types.Store[*scoringMode.ScoringMode]) *PredictionScoreResolver {
	return &PredictionScoreResolver{ir, mr}
}

func (r *PredictionScoreResolver) Resolve(ctx types.PredictionContext, bg *bettinggroup.BettingGroup) (types.PredictionScores, error) {
	scoringModeID := bg.GetScoringMode()
	scoringMode, found := r.scoringModesStore.Get(scoringModeID)

	if !found {
		// TODO: error
		return nil, errors.New("not found")
	}

	predictions, err := bg.GetPredictions(ctx.GetMatchID())

	if err != nil {
		// TODO: improve
		return nil, errors.New("error at getting predictions for match X")
	}

	return scoringMode.Resolve(ctx, predictions, bg.GetCriteria(), r.scoringItemsRepo)
}
