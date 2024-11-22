package scoring

type itemBridge struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Expression  map[string]interface{}
}

func bridgeFromScoreItem(source Item) *itemBridge {
	return &itemBridge{
		ID:          source.id,
		Name:        source.name,
		Description: source.description,
		Expression:  source.expression,
	}
}

func (b *itemBridge) toScoringItem() *Item {
	return New(b.ID, b.Name, b.Description, b.Expression)
}
