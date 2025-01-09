package scoring

import (
	"encoding/json"
	"fmt"

	"github.com/duhnnie/jexp"
	"github.com/duhnnie/jexp/expression"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type Item struct {
	id          string
	name        string
	description string
	expression  expression.Expression[bool]
}

func (i *Item) UnmarshalJSON(data []byte) error {
	var jsonStruct struct {
		ID          string          `json:"id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Expression  json.RawMessage `json:"expression"`
	}

	if err := json.Unmarshal(data, &jsonStruct); err != nil {
		return err
	}

	exp, errPath, err := jexp.New[bool](jsonStruct.Expression)

	if err != nil {
		// TODO: return custom error in which errorPath is used.
		return fmt.Errorf("error at %s: %s", errPath, err)
	}

	i.id = jsonStruct.ID
	i.name = jsonStruct.Name
	i.description = jsonStruct.Description
	i.expression = exp

	return nil
}

func (i *Item) Resolve(variables types.VariableContainer) (bool, error) {
	res, errPath, err := i.expression.Resolve(variables)

	if err != nil {
		// TODO: use custom error using errPath in it.
		return false, fmt.Errorf("error at %s: %s", errPath, err)
	}

	return res, nil
}
