package repo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duhnnie/godash"
	"github.com/duhnnie/soccerclub-scoring/constraint"
	"github.com/duhnnie/soccerclub-scoring/scoringMode"
)

type ScoringModeRepo struct{}

type strategyJSON struct {
	Type       string   `json:"type"`
	SkipIfHit  bool     `json:"skipIfThereIsScore"`
	ScoreItems []string `json:"scoreItems"`
}

type scoringModeJSON struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Strategy []strategyJSON `json:"strategy"`
}

func parseStrategy(def strategyJSON, _ int, _ []strategyJSON) scoringMode.ScoringStrategy {
	var res scoringMode.ScoringStrategy

	switch def.Type {
	case "sumFirstHit":
		res = scoringMode.NewSumFirstStrategy(def.SkipIfHit, def.ScoreItems)
	case "sumAll":
		res = scoringMode.NewSumAllStrategy(def.SkipIfHit, def.ScoreItems)
	}

	return res
}

func getData() ([]scoringModeJSON, error) {
	var dataArr []scoringModeJSON

	data, err := os.ReadFile("./json/scoring-modes.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &dataArr)

	if err != nil {
		return nil, err
	}

	return dataArr, nil
}

func getFromJSON(def scoringModeJSON) *scoringMode.ScoringMode {
	return &scoringMode.ScoringMode{
		ID:          def.ID,
		Name:        def.Name,
		Description: "",
		Strategy:    godash.Map(def.Strategy, parseStrategy),
		// TODO: complete constraints
		Constraints: []constraint.Constraint{},
	}
}

func (r *ScoringModeRepo) Get(id string) (*scoringMode.ScoringMode, error) {
	dataArr, err := getData()

	if err != nil {
		return nil, err
	}

	def, found := godash.Find(dataArr, func(item scoringModeJSON, _ int, _ []scoringModeJSON) bool {
		return item.ID == id
	})

	if !found {
		return nil, fmt.Errorf("not found")
	}

	return getFromJSON(def), nil
}

func GetAll() ([]*scoringMode.ScoringMode, error) {
	dataArr, err := getData()

	if err != nil {
		return nil, err
	}

	return godash.Map(dataArr, func(def scoringModeJSON, _ int, _ []scoringModeJSON) *scoringMode.ScoringMode {
		return getFromJSON(def)
	}), nil
}
