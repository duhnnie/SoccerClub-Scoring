package repo

import (
	"encoding/json"
	"os"

	"github.com/duhnnie/godash"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type predictionJSON struct {
	ID          string          `json:"id"`
	PredictorID string          `json:"name"`
	Body        json.RawMessage `json:"prediction"`
}

type predictionRepo struct{}

var predictionRepoInstance *predictionRepo

func PredictionRepo() *predictionRepo {
	if predictionRepoInstance == nil {
		predictionRepoInstance = &predictionRepo{}
	}

	return predictionRepoInstance
}

func (r *predictionRepo) getData() ([]predictionJSON, error) {
	var dataArr []predictionJSON

	data, err := os.ReadFile("./json/predictions.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &dataArr)

	if err != nil {
		return nil, err
	}

	return dataArr, nil
}

func (r *predictionRepo) getFromJSON(def predictionJSON) *types.Prediction {
	return types.NewPrediction(def.ID, def.PredictorID, def.Body)
}

func (r *predictionRepo) GetAll(matchID string) ([]*types.Prediction, error) {
	dataArr, err := r.getData()

	if err != nil {
		return nil, err
	}

	return godash.Map(dataArr, func(def predictionJSON, _ int, _ []predictionJSON) *types.Prediction {
		return r.getFromJSON(def)
	}), nil
}
