package scoring

import "fmt"

type Error string

func (e Error) Error() string {
	return string(e)
}

type ErrorNoScoringItemFound string

func (e ErrorNoScoringItemFound) Error() string {
	return fmt.Sprintf("no \"%s\" scoring item found", string(e))
}

const (
	NoBooleanOperationExpression = Error("expression for scoringItem needs to be of type \"booleanOperation\"")
)
