package scoring

import (
	"encoding/json"

	"github.com/duhnnie/jexp"
	"github.com/duhnnie/jexp/expression"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type JExpItem struct {
	id          string
	name        string
	description string
	expression  expression.Expression[bool]
}

func NewJExpItem(id, name, description string, expression json.RawMessage) (*JExpItem, error) {
	exp, errPath, err := jexp.New[bool](expression)

	if err != nil {
		return nil, &JExpItemError{errPath, err}
	}

	return &JExpItem{
		id:          id,
		name:        name,
		description: description,
		expression:  exp,
	}, nil
}

func (i *JExpItem) GetID() string {
	return i.id
}

func (i *JExpItem) Resolve(context types.VariableContainer) (bool, error) {
	res, errPath, err := i.expression.Resolve(context)

	if err != nil {
		return false, &JExpItemError{errPath, err}
	}

	return res, nil
}
