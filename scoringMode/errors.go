package scoringMode

import "fmt"

// TODO: change the name of "Step", it's too generic
type ErrorUknownScoringStepMode string

func (e ErrorUknownScoringStepMode) Error() string {
	return fmt.Sprintf("unknown scoringStep mode: %s", string(e))
}

type ErrorNoPointsForCriteria string

func (e ErrorNoPointsForCriteria) Error() string {
	return fmt.Sprintf("no points for criteria: %s", string(e))
}

type ErrorUknownScoringMode string

func (e ErrorUknownScoringMode) Error() string {
	return fmt.Sprintf("can't find scoring mode with id: %s", string(e))
}
