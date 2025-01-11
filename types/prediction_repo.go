package types

type PredictionRepo interface {
	GetAll(matchID string) ([]*Prediction, error)
}
