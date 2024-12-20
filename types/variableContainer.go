package types

// TODO: check if related errors should be defined here
type VariableContainer interface {
	Set(name string, data []byte) error
	GetFloat64(variableName string) (float64, error)
	GetBool(variableName string) (bool, error)
	GetString(variableName string) (string, error)
}
