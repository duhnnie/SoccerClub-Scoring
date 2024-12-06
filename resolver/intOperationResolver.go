package resolver

import (
	"math"

	"github.com/duhnnie/soccerclub-scoring/expression"
)

const (
	CLAMP        = "clamp"
	SUBSTRACTION = "substraction"
)

type intOperationResolver struct {
	ExpressionResolver *expressionResolver
}

func (r *intOperationResolver) clamp(opValue, opMin, opMax expression.Expression) (uint64, error) {
	var value, min, max uint64

	if err := r.ExpressionResolver.ResolveToUInt(opValue, &value); err != nil {
		return 0, err
	} else if err := r.ExpressionResolver.ResolveToUInt(opMin, &min); err != nil {
		return 0, err
	} else if err := r.ExpressionResolver.ResolveToUInt(opMax, &max); err != nil {
		return 0, err
	}

	clamped := math.Min(
		math.Max(
			float64(value),
			float64(min),
		),
		float64(max),
	)

	return uint64(clamped), nil
}

func (r *intOperationResolver) substract(op1, op2 expression.Expression) (uint64, error) {
	var v1, v2 uint64

	if err := r.ExpressionResolver.ResolveToUInt(op1, &v1); err != nil {
		return 0, err
	} else if err := r.ExpressionResolver.ResolveToUInt(op2, &v2); err != nil {
		return 0, err
	}

	return v1 - v2, nil
}

func (r *intOperationResolver) Resolve(name string, operands []expression.Expression) (uint64, error) {
	if name == CLAMP && len(operands) != 3 {
		return 0, &ErrorInvalidOperandsCount{name, 3, 3}
	} else if len(operands) != 2 {
		return 0, &ErrorInvalidOperandsCount{name, 2, 2}
	}

	switch name {
	case CLAMP:
		return r.clamp(operands[0], operands[1], operands[3])
	case SUBSTRACTION:
		return r.substract(operands[0], operands[1])
	default:
		return 0, ErrorUnknownOperationName(name)
	}
}
