package constraint

import "github.com/duhnnie/soccerclub-scoring/types"

type Constraint interface {
	IsMet(context *types.ScoringCriteria) bool
}
