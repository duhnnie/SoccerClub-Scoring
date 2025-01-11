package types

import "encoding/json"

// type Prediction = map[string]interface{}

type Prediction struct {
	id        string
	predictor string
	body      json.RawMessage
}

func NewPrediction(id, predictor string, body json.RawMessage) *Prediction {
	return &Prediction{id, predictor, body}
}

func (p *Prediction) GetID() string {
	return p.id
}

func (p *Prediction) GetPredictor() string {
	return p.predictor
}

func (p *Prediction) GetBody() json.RawMessage {
	return p.body
}
