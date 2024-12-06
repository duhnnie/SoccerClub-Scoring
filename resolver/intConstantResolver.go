package resolver

import "github.com/duhnnie/soccerclub-scoring/expression"

type IntConstantResolver struct{}

func (r *IntConstantResolver) Resolve(e *expression.ConstantExpression[int]) (uint64, error) {
	return uint64(e.Value), nil
}
