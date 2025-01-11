package types

// TODO: check if related errors should be defined here
type VariableContainer interface {
	Get(name string) (interface{}, error)
	ToMap() (map[string]interface{}, error)
}

type PredictionContext interface {
	VariableContainer
	SetPrediction(data []byte) error
}
