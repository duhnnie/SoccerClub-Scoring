package scoringMode

import (
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type ScoringStrategy interface {
	GetID() string
	GetName() string
	GetDescription() string
	SkipIfHit() bool
	Execute(scoreItems *scoring.Repository, context types.VariableContainer, criteria *types.ScoringCriteria) (res []*types.PredictionHit, err error)
}
