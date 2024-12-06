package scoringType

const (
	ScoringStepTypeSumFirstHit = "sumFirstHit"
	ScoringStepTypeSumAll      = "sumAll"
)

type ScoringStep struct {
	Type        string        `json:"type"`
	SkipIfScore bool          `json:"skipIfScore"`
	Items       []string      `json:"scoreItems"`
	Constraints []interface{} // TODO: implement constraints
}
