package expression

import "fmt"

type Error string

func (e Error) Error() string {
	return string(e)
}

type MissingPropertyError struct {
	expType  string
	property string
}

func (e *MissingPropertyError) Error() string {
	return fmt.Sprintf("no \"%s\" property found for \"%s\" expression type", e.property, e.expType)
}

type InvalidPropertyTypeError struct {
	expType      string
	property     string
	expectedType string
}

func (e *InvalidPropertyTypeError) Error() string {
	return fmt.Sprintf("invalid type for property \"%s\" for \"%s\" expression type: \"%s\" expected", e.property, e.expType, e.expectedType)
}

type UnknownExpressionType string

func (e UnknownExpressionType) Error() string {
	return fmt.Sprintf("unknown expression type \"%s\"", string(e))
}

const (
	ErrorInvalidExpressionType = Error("invalid data type for expression \"type\" property")
	ErrorNoExpressionTypeFound = Error("no \"type\" property found for operation expression")
)
