package scoring

import (
	"encoding/json"

	"github.com/duhnnie/jexp/expression"
	"github.com/duhnnie/soccerclub-scoring/types"
)

type Repository struct {
	items map[string]*Item
}

func NewRepository(items map[string]*Item) *Repository {
	return &Repository{
		items: items,
	}
}

func NewRepositoryFromData(data []json.RawMessage) (*Repository, int, error) {
	items := map[string]*Item{}

	for index, itemData := range data {
		var item Item

		if err := json.Unmarshal(itemData, &item); err != nil {
			return nil, index, &Error{ErrorCodeCantParseToItem, err}
		}

		items[item.id] = &item
	}

	return &Repository{
		items: items,
	}, 0, nil
}

func (r *Repository) RegisterItem(id, name, description string, expression expression.Expression[bool]) {
	r.items[id] = &Item{
		id:          id,
		name:        name,
		description: description,
		expression:  expression,
	}
}

func (r *Repository) ExecuteItem(itemID string, variables types.VariableContainer) (bool, error) {
	if scoringItem, exists := r.items[itemID]; !exists {
		return false, ErrorNoScoringItemFound(itemID)
	} else {
		return scoringItem.Resolve(variables)
	}
}
