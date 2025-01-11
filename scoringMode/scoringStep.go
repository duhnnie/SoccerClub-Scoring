package scoringMode

const (
	ScoringStepTypeSumFirstHit = "sumFirstHit"
	ScoringStepTypeSumAll      = "sumAll"
)

// TODO: Remove
type ScoringStep struct {
	Type        string        `json:"type"`
	SkipIfScore bool          `json:"skipIfScore"`
	Items       []string      `json:"scoreItems"`
	Constraints []interface{} // TODO: implement constraints
}
