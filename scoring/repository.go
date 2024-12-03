package scoring

import (
	"encoding/json"
	"fmt"
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

func (r *Repository) ExecuteItem(itemID string) (bool, error) {
	if scoringItem, exists := r.items[itemID]; !exists {
		return false, NoScoringItemFound
	} else {
		// TODO: call item.resolve()
		fmt.Printf("Expression: %+v", scoringItem.expression)
		return true, nil
	}
}
