package expression

type operationExpressionType string
type variableExpressionType string
type constantExpressionType string

type constantExpressionValueType interface {
	int | bool
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
