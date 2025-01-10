package scoring

import (
	"github.com/duhnnie/godash"
	"github.com/duhnnie/soccerclub-scoring/types"
	jsone "github.com/json-e/json-e/v4"
)

type JsoneItem struct {
	id          string
	name        string
	description string
	expression  string
}

func NewJSONeItem(id, name, description, expression string) *JsoneItem {
	return &JsoneItem{id, name, description, expression}
}

func (i *JsoneItem) GetID() string {
	return i.id
}

func (i *JsoneItem) Resolve(variables types.VariableContainer) (bool, error) {
	var template = map[string]interface{}{
		"$eval": i.expression,
	}

	match, err := variables.Get("match")

	if err != nil {
		return false, err
	}

	prediction, err := variables.Get("prediction")

	if err != nil {
		return false, err
	}

	var context = map[string]interface{}{
		"match":      match,
		"prediction": prediction,
		"clamp":      jsone.WrapFunction(func(a, b, c float64) float64 { return godash.Clamp(a, b, c) }),
	}

	value, err := jsone.Render(template, context)

	if err != nil {
		return false, err
	}

	if boolValue, ok := value.(bool); !ok {
		return false, DoesntResolveToBooleanError(i.expression)
	} else {
		return boolValue, nil
	}
}
