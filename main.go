package main

import (
	"fmt"

	"github.com/duhnnie/soccerclub-scoring/scoringItem"
)

func main() {
	// data, err := os.ReadFile("./json/scoring-items.json")

	json := `
		{
    "name": "Score Hit",
    "description": "When prediction hit exact score",
    "expression": {
      "type": "booleanOperation",
      "name": "and",
      "operands": [
        {
          "type": "booleanOperation",
          "name": "eq",
          "operands": [
            {
              "type": "intVariable",
              "value": "match.home.score"
            },
            {
              "type": "intVariable",
              "value": "prediction.home.score"
            }
          ]
        },
        {
          "type": "booleanOperation",
          "name": "eq",
          "operands": [
            {
              "type": "intVariable",
              "value": "match.away.score"
            },
            {
              "type": "intVariable",
              "value": "prediction.away.score"
            }
          ]
        }
      ]
    }
  }
	`

	si, err := scoringItem.NewFromString("score-hit", json)

	fmt.Println(si, err)

	fmt.Println(si.GetID())
	fmt.Println(si.GetName())
	fmt.Println(si.GetDescription())
	// fmt.Println(si.GetExpression())

	// fmt.Println(si.ID)
	// fmt.Println(si.Name)
	// fmt.Println(si.Description)
	// fmt.Println(si.Expression)

	data, _ := si.ToJSON()

	fmt.Println(string(data))
}
