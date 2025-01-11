package scoring

import (
	"encoding/json"

	"github.com/duhnnie/jexp/expression"
)

type Repository struct {
	items map[string]Item
}

func NewRepository(items map[string]Item) *Repository {
	return &Repository{
		items: items,
	}
}

func (r *Repository) UnmarshalJSON(data []byte) error {
	var itemDefs []struct {
		ID               string          `json:"id"`
		Name             string          `json:"name"`
		Description      string          `json:"description"`
		ExpressionEngine string          `json:"expEngine"`
		Expression       json.RawMessage `json:"expression"`
	}

	if err := json.Unmarshal(data, &itemDefs); err != nil {
		return err
	}

	items := map[string]Item{}
	var item Item
	var err error

	for index, def := range itemDefs {
		switch def.ExpressionEngine {
		case "jexp":
			item, err = NewJExpItem(def.ID, def.Name, def.Description, def.Expression)
		case "jsone":
			var expression string

			err = json.Unmarshal(def.Expression, &expression)
			item = NewJSONeItem(def.ID, def.Name, def.Description, expression)
		default:
			return NotSupporteEngineError(def.ExpressionEngine)
		}

		if err != nil {
			return &UnmarshalRepositoryError{index, err}
		}

		items[item.GetID()] = item
	}

	r.items = items

	return nil
}

func (r *Repository) RegisterItem(id, name, description string, expression expression.Expression[bool]) {
	r.items[id] = &JExpItem{
		id:          id,
		name:        name,
		description: description,
		expression:  expression,
	}
}

func (r *Repository) Get(id string) Item {
	return r.items[id]
}

// TODO: Remove this method, repository should execute anything
// func (r *Repository) ExecuteItem(itemID string, context types.VariableContainer) (bool, error) {
// 	if scoringItem, exists := r.items[itemID]; !exists {
// 		return false, ErrorNoScoringItemFound(itemID)
// 	} else {
// 		return scoringItem.Resolve(context)
// 	}
// }
