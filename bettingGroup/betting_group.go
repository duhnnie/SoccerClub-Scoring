package bettinggroup

// TODO: check convnetion for package names

import "github.com/duhnnie/soccerclub-scoring/types"

type BettingGroup struct {
	id             string
	criteria       *types.ScoringCriteria
	scoringMode    string
	predictionRepo types.PredictionRepo
}

func New(id string, criteria *types.ScoringCriteria, scoringMode string, predictionsRepo types.PredictionRepo) *BettingGroup {
	return &BettingGroup{id, criteria, scoringMode, predictionsRepo}
}

func (bg *BettingGroup) GetCriteria() *types.ScoringCriteria {
	return bg.criteria
}

func (bg *BettingGroup) GetScoringMode() string {
	return bg.scoringMode
}

func (bg *BettingGroup) GetPredictions(matchID string) ([]*types.Prediction, error) {
	return bg.predictionRepo.GetAll(matchID)
}
