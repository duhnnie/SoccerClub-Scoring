package variable

import "fmt"

type Error string

func (err Error) Error() string {
	return string(err)
}

type ErrorNoVariableFound string

func (e ErrorNoVariableFound) Error() string {
	return fmt.Sprintf("no \"%s\" variable was found", string(e))
}

type ErrorCantResolveToType string

func (e ErrorCantResolveToType) Error() string {
	return fmt.Sprintf("can't resolve to type: %s", string(e))
}

const (
	ErrorResolveInvalidParams     = Error("second argument needs to be a string or a slice of strings.")
	ErrorResolveInvalidFirstParam = Error("\"target\" parameter should be a \"map[string]interface{}\" when second argument is not an empty string nor empty string slice")
)
