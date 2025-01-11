package constraint

import (
	"errors"

	"github.com/duhnnie/soccerclub-scoring/types"
	jsone "github.com/json-e/json-e/v4"
)

type JSONeConstraint struct {
	expression string
}

func NewJSONe(expression string) *JSONeConstraint {
	return &JSONeConstraint{expression}
}

func (c *JSONeConstraint) IsMet(context types.ScoringCriteria) (bool, error) {
	var template = map[string]interface{}{
		"$eval": c.expression,
	}

	value, err := jsone.Render(template, context.TopMap())

	if err != nil {
		return false, err
	}

	if boolValue, ok := value.(bool); !ok {
		// TODO: re-use error
		// return false, DoesntResolveToBooleanError(i.expression)
		return false, errors.New("it doesn't resolve to boolean")
	} else {
		return boolValue, nil
	}
}
