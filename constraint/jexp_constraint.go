package constraint

import (
	"encoding/json"

	"github.com/duhnnie/jexp"
	"github.com/duhnnie/jexp/expression"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type JExpConstraint struct {
	expression expression.Expression[bool]
}

func NewJExp(expression json.RawMessage) (*JExpConstraint, error) {
	exp, _, err := jexp.New[bool](expression)

	if err != nil {
		// TODO: use an error that uses errPath
		// return nil, &JExpItemError{errPath, err}
		return nil, err
	}

	return &JExpConstraint{expression: exp}, nil
}

func (c *JExpConstraint) IsMet(context *types.ScoringCriteria) (bool, error) {
	res, _, err := c.expression.Resolve(context)

	if err != nil {
		// TODO: use error with errPath
		// return false, &JExpItemError{errPath, err}
		return false, err
	}

	return res, nil
}
