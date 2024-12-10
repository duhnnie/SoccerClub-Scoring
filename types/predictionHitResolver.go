package types

type PredictionHitResolver interface {
	Resolve(scoringModeID string, vars VariableContainer, criteria ScoringCriteria) ([]*PredictionHit, error)
}
