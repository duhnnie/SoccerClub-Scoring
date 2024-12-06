package scoringType

import "fmt"

// TODO: change the name of "Step", it's too generic
type ErrorUknownScoringStepType string

func (e ErrorUknownScoringStepType) Error() string {
	return fmt.Sprintf("unknown scoringStep type: %s", string(e))
}

type ErrorNoPointsForSCriteria string

func (e ErrorNoPointsForSCriteria) Error() string {
	return fmt.Sprintf("no point for criteria: %s", string(e))
}

type ErrorUknownScoringType string

func (e ErrorUknownScoringType) Error() string {
	return fmt.Sprintf("can't find scoring type with id: %s", string(e))
}
