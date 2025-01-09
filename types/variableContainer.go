package types

// TODO: check if related errors should be defined here
type VariableContainer interface {
	Get(name string) (interface{}, error)
	// Set(name string, data []byte) error
	// GetFloat64(variableName string) (float64, error)
	// GetBool(variableName string) (bool, error)
	// GetString(variableName string) (string, error)
}
