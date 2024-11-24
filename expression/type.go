package expression

type operationExpressionType string
type variableExpressionType string
type constantExpressionType string

type constantExpressionValueType interface {
	int | bool
}

const (
	ExpressionTypeBooleanOperation = operationExpressionType("booleanOperation")
	ExpressionTypeIntOperation     = operationExpressionType("intOperation")
	ExpressionTypeIntVariable      = variableExpressionType("intVariable")
	ExpressionTypeIntConstant      = constantExpressionType("intConstant")
)

type Expression interface {
	GetType() string
}

type OperationExpression struct {
	expressionExpType[operationExpressionType]
	name     string
	operands []Expression
}

type VariableExpression struct {
	expressionExpType[variableExpressionType]
	name string
}

type ConstantExpression[U constantExpressionValueType] struct {
	expressionExpType[constantExpressionType]
	value U
}
