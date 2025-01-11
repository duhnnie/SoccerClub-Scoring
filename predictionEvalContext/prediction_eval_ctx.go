package predictionevalcontext

import (
	"github.com/duhnnie/soccerclub-scoring/match"
	"github.com/duhnnie/soccerclub-scoring/types"
	"github.com/duhnnie/valuebox"
)

type PredictionEvalCtx struct {
	matchID   string
	container valuebox.Box
}

func NewPredictionEvalCtx() *PredictionEvalCtx {
	return &PredictionEvalCtx{
		container: *valuebox.New(),
	}
}

func (c *PredictionEvalCtx) SetMatch(m *match.Match) {
	c.matchID = m.GetID()
	c.container.Set("match", m.GetFacts())
}

func (c *PredictionEvalCtx) GetMatchID() string {
	return c.matchID
}

func (c *PredictionEvalCtx) SetPrediction(p *types.Prediction) {
	c.container.Set("prediction", p.GetBody())
}

func (c *PredictionEvalCtx) Get(name string) (interface{}, error) {
	return c.container.Get(name)
}
