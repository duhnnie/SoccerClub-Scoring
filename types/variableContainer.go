package types

// TODO: check if related errors should be defined here
type VariableContainer interface {
	Get(name string) (interface{}, error)
}
