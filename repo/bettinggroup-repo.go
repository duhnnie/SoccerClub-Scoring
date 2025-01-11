package repo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duhnnie/godash"
	bettinggroup "github.com/duhnnie/soccerclub-scoring/bettingGroup"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type bettingGroupJSON struct {
	ID            string             `json:"id"`
	Criteria      map[string]float64 `json:"criteria"`
	ScoringModeID string             `json:"scoringMode"`
}

type bettingGroupRepo struct {
}

var bettingGroupRepoInstance *bettingGroupRepo

func BettingGroupRepo() *bettingGroupRepo {
	if bettingGroupRepoInstance == nil {
		bettingGroupRepoInstance = &bettingGroupRepo{}
	}

	return bettingGroupRepoInstance
}

func (r *bettingGroupRepo) getData() ([]bettingGroupJSON, error) {
	var dataArr []bettingGroupJSON

	data, err := os.ReadFile("./json/betting-group.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &dataArr)

	if err != nil {
		return nil, err
	}

	return dataArr, nil
}

func (r *bettingGroupRepo) getFromJSON(def bettingGroupJSON) *bettinggroup.BettingGroup {
	sc := types.NewScoringCriteria()

	for k, v := range def.Criteria {
		sc.Set(k, v)
	}

	return bettinggroup.New(def.ID, sc, def.ScoringModeID, PredictionRepo())
}

func (r *bettingGroupRepo) Get(id string) (*bettinggroup.BettingGroup, error) {
	dataArr, err := r.getData()

	if err != nil {
		return nil, err
	}

	def, found := godash.Find(dataArr, func(item bettingGroupJSON, _ int, _ []bettingGroupJSON) bool {
		return item.ID == id
	})

	if !found {
		return nil, fmt.Errorf("not found")
	}

	return r.getFromJSON(def), nil
}
