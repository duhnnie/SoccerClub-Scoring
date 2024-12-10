package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/scoringMode"
	"github.com/duhnnie/soccerclub-scoring/variable"
)

type StringType string

type ToMarshal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (t ToMarshal) MarshalJSON() ([]byte, error) {
	return []byte("{}"), nil
}

func (t ToMarshal) UnmarshalJSON(data []byte) error {
	return nil
}

// func main() {
// 	data, _ := os.ReadFile("./json/scoring-items.json")
// 	repo, err := scoring.NewRepositoryFromData(data)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	_, err = repo.ExecuteItem("score-hitsd", variable.NewRepo())

// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		println("good")
// 	}

// 	matchData, _ := os.ReadFile("./json/match.json")
// 	var matchVars map[string]interface{}
// 	_ = json.Unmarshal(matchData, &matchVars)

// 	predictionData, _ := os.ReadFile("./json/prediction.json")
// 	var predictionVars map[string]interface{}
// 	_ = json.Unmarshal(predictionData, &predictionVars)

// 	vars := variable.NewRepo()
// 	vars.Set("match", matchVars)
// 	vars.Set("prediction", predictionVars)

// 	v, err := repo.ExecuteItem("one-side-score-hit", vars)

// 	fmt.Println(v, err)
// }

func main() {
	data, _ := os.ReadFile("./json/scoring-items.json")
	repo, err := scoring.NewRepositoryFromData(data)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	matchData, _ := os.ReadFile("./json/match.json")
	var matchVars map[string]interface{}
	_ = json.Unmarshal(matchData, &matchVars)

	predictionData, _ := os.ReadFile("./json/prediction.json")
	var predictionVars map[string]interface{}
	_ = json.Unmarshal(predictionData, &predictionVars)

	vars := variable.NewRepo()
	vars.Set("match", matchVars)
	vars.Set("prediction", predictionVars)

	data, _ = os.ReadFile("./json/scoring-types.json")
	scoringModeRepo, err := scoringMode.NewRepoFromData(data, repo)
	if err != nil {
		panic(err)
	}

	criteria := map[string]float64{
		"score-hit":          8,
		"score-diff-hit":     5,
		"winner-hit":         3,
		"one-side-score-hit": 1,
	}

	scores, err := scoringModeRepo.Resolve("all-hits", vars, criteria)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, score := range scores {
		fmt.Printf("%+v\n", score)
	}

	// fmt.Printf("%+v\n", scores)
}
