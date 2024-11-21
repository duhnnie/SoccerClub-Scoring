package main

import (
	"fmt"
	"os"

	"github.com/duhnnie/soccerclub-scoring/scoring"
)

func main() {
	data, _ := os.ReadFile("./json/scoring-items.json")
	repo, _ := scoring.NewItemRepositoryFromData(data)
	_, err := repo.ExecuteItem("score-hit")

	if err != nil {
		fmt.Println("error", err)
	} else {
		println("good")
	}

}
