package main

import (
	"fmt"
	"os"

	"github.com/duhnnie/soccerclub-scoring/scoring"
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
