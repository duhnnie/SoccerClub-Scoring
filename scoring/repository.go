package scoring

import (
	"encoding/json"

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

func NewRepositoryFromData(data []byte) (*Repository, error) {
	var jsonStruct []json.RawMessage
	items := map[string]*Item{}

	if err := json.Unmarshal(data, &jsonStruct); err != nil {
		return nil, err
	}

	for _, itemData := range jsonStruct {
		var item Item

		if err := json.Unmarshal(itemData, &item); err != nil {
			return nil, err
		}

		items[item.id] = &item
	}

	return &Repository{
		items: items,
	}, nil
}

func (r *Repository) ExecuteItem(itemID string, variables types.VariableContainer) (bool, error) {
	if scoringItem, exists := r.items[itemID]; !exists {
		return false, ErrorNoScoringItemFound(itemID)
	} else {
		return scoringItem.Resolve(variables)
	}
}
