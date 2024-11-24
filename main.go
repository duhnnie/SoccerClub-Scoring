package main

import (
	"fmt"
	"os"

	"github.com/duhnnie/soccerclub-scoring/scoring"
	"github.com/duhnnie/soccerclub-scoring/variable"
)

func main() {
	data, _ := os.ReadFile("./json/scoring-items.json")
	repo, _ := scoring.NewRepositoryFromData(data)
	_, err := repo.ExecuteItem("score-hit")

	if err != nil {
		fmt.Println("error", err)
	} else {
		println("good")
	}

	vars := make(map[string]interface{})
	vars["home"] = make(map[string]interface{})
	vars["away"] = make(map[string]interface{})
	vars["home"].(map[string]interface{})["score"] = 3
	vars["away"].(map[string]interface{})["score"] = 2

	varRepo := variable.NewRepo()
	varRepo.Set("match", vars)

	fmt.Println(varRepo.GetFloat32("match.home.score"))
	// fmt.Println(variable.Resolve(vars, "home.score"))

	// x := make(map[string]interface{})
	// x["asdfsad"] = 4
	// expression.Test()
}
