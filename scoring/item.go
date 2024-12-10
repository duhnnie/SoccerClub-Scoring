package scoring

import (
	"encoding/json"

	"github.com/duhnnie/soccerclub-scoring/expression"
	"github.com/duhnnie/soccerclub-scoring/resolver"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type Item struct {
	id          string
	name        string
	description string
	expression  *expression.OperationExpression
}

func (i *Item) UnmarshalJSON(data []byte) error {
	var jsonStruct struct {
		ID          string          `json:"id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Expression  json.RawMessage `json:"expression"`
	}

	if err := json.Unmarshal(data, &jsonStruct); err != nil {
		return err
	}

	exp, err := expression.UnmarshalExpression(jsonStruct.Expression)

	if err != nil {
		return err
	}

	var opExp *expression.OperationExpression
	var ok bool

	if opExp, ok = exp.(*expression.OperationExpression); !ok {
		return NoBooleanOperationExpression
	} else if opExp.Type != expression.ExpTypeBooleanOperation {
		return NoBooleanOperationExpression
	}

	i.id = jsonStruct.ID
	i.name = jsonStruct.Name
	i.description = jsonStruct.Description
	i.expression = opExp

	return nil
}

func (i *Item) Resolve(variables types.VariableContainer) (bool, error) {
	if i.expression.Type != expression.ExpTypeBooleanOperation {
		return false, NoBooleanOperationExpression
	}

	r := resolver.New(variables)
	res, err := r.Resolve(i.expression)

	if err != nil {
		return false, err
	}

	if v, ok := res.(bool); !ok {
		return false, resolver.ErrorCantResolveToType("bool")
	} else {
		return v, nil
	}
}
