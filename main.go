package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duhnnie/soccerclub-scoring/match"
	predictionevalcontext "github.com/duhnnie/soccerclub-scoring/predictionEvalContext"
	"github.com/duhnnie/soccerclub-scoring/repo"
	"github.com/duhnnie/soccerclub-scoring/resolver"
	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/store"
)

func main() {
	scoringItemsData, _ := os.ReadFile("./json/scoring-items.json")
	scoringItemsRepo := &scoring.Repository{}

	if err := json.Unmarshal(scoringItemsData, scoringItemsRepo); err != nil {
		panic(err)
	}

	s, err := store.NewScoringModeStore()

	if err != nil {
		panic(err)
	}

	bg, err := repo.BettingGroupRepo().Get("1234567")

	if err != nil {
		panic(err)
	}

	m, err := match.Repository().Get("asdf")

	if err != nil {
		panic(err)
	}

	// predictions, err := bg.GetPredictions(m.GetID())

	// if err != nil {
	// 	panic(err)
	// }

	r := resolver.New(scoringItemsRepo, s)
	ctx := &predictionevalcontext.PredictionEvalCtx{}
	ctx.SetMatch(m)

	scores, err := r.Resolve(ctx, bg)

	if err != nil {
		panic(err)
	}

	fmt.Println(scores)
}
