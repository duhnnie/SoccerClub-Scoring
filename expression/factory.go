package expression

import (
	"errors"
)

type rawExpression = map[string]interface{}

func toExpressions(defs []interface{}) ([]Expression, int, error) {
	var expressions []Expression

	for i, def := range defs {
		if def, ok := def.(rawExpression); !ok {
			return expressions, i, ErrorInvalidExpressionDefType
		} else if expression, err := Create(def); err != nil {
			return expressions, i, err
		} else {
			expressions = append(expressions, expression)
		}
	}

	return expressions, 0, nil
}

func CreateOperationExpression(expType string, def rawExpression) (*OperationExpression, error) {
	if operationName, exists := def["name"]; !exists {
		return nil, &MissingPropertyError{string(expType), "name"}
	} else if operationName, ok := operationName.(string); !ok {
		return nil, &InvalidPropertyTypeError{string(expType), "name", "string"}
	} else if operandsInterface, exists := def["operands"]; !exists {
		return nil, &MissingPropertyError{string(expType), "operands"}
	} else if operands, ok := operandsInterface.([]interface{}); !ok {
		return nil, &InvalidPropertyTypeError{expType, "operands", "[]interface{}"}
	} else {
		expressions, index, err := toExpressions(operands)

		if err != nil {
			return nil, &ChildError{
				index:   index,
				err:     err,
				expType: expType,
			}
		}

		return &OperationExpression{
			expressionExpType: expressionExpType{expType},
			name:              operationName,
			operands:          expressions,
		}, nil
	}
}

func CreateVariableExpression(expType string, def rawExpression) (*VariableExpression, error) {
	if variableName, exists := def["name"]; !exists {
		return nil, &MissingPropertyError{string(expType), "name"}
	} else if variableName, ok := variableName.(string); !ok {
		return nil, &InvalidPropertyTypeError{string(expType), "name", "string"}
	} else {
		return &VariableExpression{
			expressionExpType: expressionExpType{expType},
			name:              variableName,
		}, nil
	}
}

func CreateIntConstantExpression(def rawExpression) (*ConstantExpression[int], error) {
	expType := "intConstant"

	if value, exists := def["value"]; !exists {
		return nil, &MissingPropertyError{string(expType), "value"}
	} else if value, ok := value.(int); !ok {
		return nil, &InvalidPropertyTypeError{string(expType), "value", "int"}
	} else {
		return &ConstantExpression[int]{
			expressionExpType: expressionExpType{expType},
			value:             value,
		}, nil
	}
}

func Create(def rawExpression) (Expression, error) {
	if expTypeInterface, exists := def["type"]; !exists {
		return nil, ErrorNoExpressionTypeFound
	} else if expType, ok := expTypeInterface.(string); !ok {
		return nil, errors.New("\"type\" is not a string")
	} else {
		switch expType {
		case "intOperation", "booleanOperation":
			return CreateOperationExpression(expType, def)
		case "intVariable":
			return CreateVariableExpression(expType, def)
		case "intConstant":
			return CreateIntConstantExpression(def)
		default:
			return nil, UnknownExpressionType(expType)
		}
	}
}
