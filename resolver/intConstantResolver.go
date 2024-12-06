package resolver

import "github.com/duhnnie/soccerclub-scoring/expression"

type IntConstantResolver struct{}

func (r *IntConstantResolver) Resolve(e *expression.ConstantExpression[int]) (int64, error) {
	return int64(e.Value), nil
}
