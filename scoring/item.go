package scoring

import (
	"encoding/json"

	"github.com/duhnnie/soccerclub-scoring/expression"
)

type Item struct {
	id          string
	name        string
	description string
	expression  expression.Expression
}

func New(id, name, description string, expression expression.Expression) *Item {
	return &Item{
		id:          id,
		name:        name,
		description: description,
		expression:  expression,
	}
}

func NewFromData(id string, data []byte) (*Item, error) {
	bridge := itemBridge{}

	if err := json.Unmarshal(data, &bridge); err != nil {
		return nil, err
	}

	bridge.ID = id
	return bridge.toScoringItem()
}

func NewFromString(id, data string) (*Item, error) {
	return NewFromData(id, []byte(data))
}

// func (s *Item) ToJSON() ([]byte, error) {
// 	bridge := bridgeFromScoreItem(*s)
// 	return json.Marshal(bridge)
// }

func (s *Item) GetID() string {
	return s.id
}

func (s *Item) GetName() string {
	return s.name
}

func (s *Item) GetDescription() string {
	return s.description
}

func (s *Item) GetExpression() expression.Expression {
	return s.expression
}
