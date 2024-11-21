package scoringItem

type scoringItemBridge struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Expression  map[string]interface{}
}

func bridgeFromScoreItem(source ScoringItem) *scoringItemBridge {
	return &scoringItemBridge{
		ID:          source.id,
		Name:        source.name,
		Description: source.description,
		Expression:  source.expression,
	}
}

func (b *scoringItemBridge) toScoringItem(target *ScoringItem) {
	target.id = b.ID
	target.name = b.Name
	target.description = b.Description
	target.expression = b.Expression
}
