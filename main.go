package main

import (
	"encoding/json"
	"fmt"

	"github.com/duhnnie/soccerclub-scoring/expression"
)

type StringType string

func main() {
	// data, _ := os.ReadFile("./json/scoring-items.json")
	// repo, _ := scoring.NewRepositoryFromData(data)
	// _, err := repo.ExecuteItem("score-hit")

	// if err != nil {
	// 	fmt.Println("error", err)
	// } else {
	// 	println("good")
	// }

	// vars := make(map[string]interface{})
	// vars["home"] = make(map[string]interface{})
	// vars["away"] = make(map[string]interface{})
	// vars["home"].(map[string]interface{})["score"] = 3
	// vars["away"].(map[string]interface{})["score"] = 2

	// varRepo := variable.NewRepo()
	// varRepo.Set("match", vars)

	// fmt.Println(varRepo.GetInt("match.home.score"))
	// fmt.Println(variable.Resolve(vars, "home.score"))

	// --------- Testing expressions

	x := make(map[string]interface{})

	// var x struct {
	// 	Type string `json:"type"`
	// 	Name string `json:"name"`
	// 	// Operands []map[string]interface{}
	// 	Operands []expression.Expression
	// }

	jsonString := `
		{
			"type": "intOperation",
			"name": "substraction",
			"operands": [
				{
					"type": "intVariable",
					"name": "match.home.score"
				},
				{
					"type": "intVariable",
					"name": "match.away.score"
				}
			]
		}
	`

	json.Unmarshal([]byte(jsonString), &x)

	y, err := expression.Create(x)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Printf("%+v\n", y)
	fmt.Printf("%+v\n", y.GetType())

	if op, ok := y.(*expression.OperationExpression); !ok {
		fmt.Println("It's not an operation expression")
	} else if firstOperand, ok := op.GetOperands()[0].(*expression.VariableExpression); !ok {
		fmt.Println("operands property is not an array")
	} else {
		fmt.Printf("%+v\n", firstOperand)
	}

}
