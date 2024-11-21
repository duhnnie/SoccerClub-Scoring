package scoring

type scoringItemBridge struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Expression  map[string]interface{}
}

func bridgeFromScoreItem(source Item) *scoringItemBridge {
	return &scoringItemBridge{
		ID:          source.id,
		Name:        source.name,
		Description: source.description,
		Expression:  source.expression,
	}
}

func (b *scoringItemBridge) toScoringItem() *Item {
	return New(b.ID, b.Name, b.Description, b.Expression)
}
