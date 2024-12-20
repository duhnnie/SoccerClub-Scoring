package resolver

import (
	"github.com/duhnnie/soccerclub-scoring/expression"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type expressionResolver struct {
	variables         types.VariableContainer
	booleanOpResolver *booleanOperationResolver
	intOpResolver     *intOperationResolver
	intConstResolver  *IntConstantResolver
}

func New(variables types.VariableContainer) *expressionResolver {
	booleanOpResolver := booleanOperationResolver{}
	intOpResolver := intOperationResolver{}
	intConstResolver := IntConstantResolver{}

	this := expressionResolver{
		variables:         variables,
		booleanOpResolver: &booleanOpResolver,
		intOpResolver:     &intOpResolver,
		intConstResolver:  &intConstResolver,
	}

	booleanOpResolver.ExpressionResolver = &this
	intOpResolver.ExpressionResolver = &this

	return &this
}

func (r *expressionResolver) resolveOperationExpression(exp *expression.OperationExpression) (interface{}, error) {
	switch exp.Type {
	case expression.ExpTypeBooleanOperation:
		return r.booleanOpResolver.Resolve(exp.Name, exp.Operands)
	case expression.ExpTypeIntOperation:
		return r.intOpResolver.Resolve(exp.Name, exp.Operands)
	default:
		return nil, expression.ErrorUnknownExpressionType(exp.Type)
	}
}

func (r *expressionResolver) resolveVariableExpression(exp *expression.VariableExpression) (interface{}, error) {
	switch exp.Type {
	case expression.ExpTypeIntVariable:
		if v, err := r.variables.GetFloat64(exp.Name); err != nil {
			return 0, err
		} else {
			return int64(v), nil
		}
	default:
		return nil, expression.ErrorUnknownExpressionType(exp.Type)
	}
}

func (r *expressionResolver) Resolve(v any) (interface{}, error) {
	switch exp := v.(type) {
	case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, string, bool:
		return v, nil
	case *expression.OperationExpression:
		return r.resolveOperationExpression(v.(*expression.OperationExpression))
	case *expression.VariableExpression:
		return r.resolveVariableExpression(v.(*expression.VariableExpression))
	case *expression.ConstantExpression[int]:
		return r.intConstResolver.Resolve(exp)
	default:
		return nil, ErrorCantResolveToExpression
	}
}

func (r *expressionResolver) ResolveToBoolean(e expression.Expression, out *bool) error {
	if res, err := r.Resolve(e); err != nil {
		return err
	} else if v, ok := res.(bool); ok {
		*out = v
		return nil
	} else {
		return ErrorCantResolveToType("bool")
	}
}

func (r *expressionResolver) ResolveToInt(e expression.Expression, out *int64) error {
	if res, err := r.Resolve(e); err != nil {
		return err
	} else if v, ok := res.(int64); ok {
		*out = v
		return nil
	} else {
		return ErrorCantResolveToType("int64")
	}
}

func (r *expressionResolver) ResolveToUInt(e expression.Expression, out *uint64) error {
	if res, err := r.Resolve(e); err != nil {
		return err
	} else if v, ok := res.(uint64); ok {
		*out = v
		return nil
	}

	return ErrorCantResolveToType("uint64")
}
