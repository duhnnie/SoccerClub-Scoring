package scoring

import "github.com/duhnnie/soccerclub-scoring/expression"

type itemBridge struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Expression  map[string]interface{}
}

// func bridgeFromScoreItem(source Item) *itemBridge {
// 	return &itemBridge{
// 		ID:          source.id,
// 		Name:        source.name,
// 		Description: source.description,
// 		Expression:  source.expression,
// 	}
// }

func (b *itemBridge) toScoringItem() (*Item, error) {
	exp, err := expression.Create(b.Expression)

	if err != nil {
		return nil, err
	}

	return New(b.ID, b.Name, b.Description, exp), nil
}
