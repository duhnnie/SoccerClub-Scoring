package expression

type rawExpression map[string]interface{}

func CreateOperationExpression(expType operationExpressionType, def rawExpression) (*OperationExpression, error) {
	if operationName, exists := def["name"]; !exists {
		return nil, &MissingPropertyError{string(expType), "name"}
	} else if operationName, ok := operationName.(string); !ok {
		return nil, &InvalidPropertyTypeError{string(expType), "name", "string"}
	} else {
		return &OperationExpression{
			expressionExpType: expressionExpType[operationExpressionType]{expType},
			name:              operationName,
			operands:          []Expression{},
		}, nil
	}

}

func CreateVariableExpression(expType variableExpressionType, def rawExpression) (*VariableExpression, error) {
	if variableName, exists := def["name"]; !exists {
		return nil, &MissingPropertyError{string(expType), "name"}
	} else if variableName, ok := variableName.(string); !ok {
		return nil, &InvalidPropertyTypeError{string(expType), "name", "string"}
	} else {
		return &VariableExpression{
			expressionExpType: expressionExpType[variableExpressionType]{expType},
			name:              variableName,
		}, nil
	}
}

func CreateIntConstantExpression(def rawExpression) (*ConstantExpression[int], error) {
	expType := ExpressionTypeIntConstant

	if value, exists := def["value"]; !exists {
		return nil, &MissingPropertyError{string(expType), "value"}
	} else if value, ok := value.(int); !ok {
		return nil, &InvalidPropertyTypeError{string(expType), "value", "int"}
	} else {
		return &ConstantExpression[int]{
			expressionExpType: expressionExpType[constantExpressionType]{expType},
			value:             value,
		}, nil
	}
}

func Create(def rawExpression) (Expression, error) {
	if expType, exists := def["type"]; !exists {
		return nil, ErrorNoExpressionTypeFound
	} else if operationExpType, ok := expType.(operationExpressionType); ok {
		return CreateOperationExpression(operationExpType, def)
	} else if variableExpType, ok := expType.(variableExpressionType); ok {
		return CreateVariableExpression(variableExpType, def)
	} else if constantExpType, ok := expType.(constantExpressionType); ok {
		switch constantExpType {
		case ExpressionTypeIntConstant:
			return CreateIntConstantExpression(def)
		default:
			return nil, UnknownExpressionType(constantExpType)
		}
	} else if expTypeString, ok := expType.(string); !ok {
		return nil, UnknownExpressionType(expTypeString)
	}

	return nil, ErrorInvalidExpressionType
}
