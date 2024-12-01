package expression

type constantExpressionValueType interface {
	int | float64 | bool
}

type Expression interface {
	GetType() string
}

type OperationExpression struct {
	expressionExpType
	name     string
	operands []Expression
}

func (exp *OperationExpression) GetOperands() []Expression {
	return exp.operands
}

type VariableExpression struct {
	expressionExpType
	name string
}

type ConstantExpression[U constantExpressionValueType] struct {
	expressionExpType
	value U
}
