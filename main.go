package main

import (
	"fmt"
	"os"

	"github.com/duhnnie/soccerclub-scoring/scoring"
)

type StringType string

func main() {
	data, _ := os.ReadFile("./json/scoring-items.json")
	repo, err := scoring.NewRepositoryFromData(data)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = repo.ExecuteItem("score-hit")

	if err != nil {
		fmt.Println("error", err)
	} else {
		println("good")
	}
}
