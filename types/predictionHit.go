package types

import "fmt"

// TODO: Maybe this is not gonna be needed anymore
type PredictionHit struct {
	ScoringItem string
	Points      float64
}

func (ph *PredictionHit) String() string {
	return fmt.Sprintf("%s:%.2f", ph.ScoringItem, ph.Points)
}
